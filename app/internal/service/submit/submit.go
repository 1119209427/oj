package submit

import (
	"bytes"
	"fmt"
	"io"
	"oj/app/define"
	g "oj/app/global"
	"oj/app/internal/model"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type sSubmit struct {
}

var (
	onceSubmit sync.Once
	insSubmit  *sSubmit
)

func NewSubmitService() *sSubmit {
	onceSubmit.Do(func() {
		insSubmit = &sSubmit{}
	})
	return insSubmit
}

func (s *sSubmit) GetProblemAndSubmitIdByUserId(uid, page, size int) ([]string, error) {
	return s.getProblemAndSubmitIdByRedis(uid, page, size)
}

func (s *sSubmit) getProblemAndSubmitIdByRedis(uid, page, size int) ([]string, error) {
	//第一步:查询key为uid，是否存在，存在就直接返回，不存在则去数据库查询
	uidString := strconv.Itoa(uid)
	if n, err := g.DbSubmit.Exists(g.RedisContext, uidString).Result(); n > 0 {
		//存在
		if err != nil {
			g.Logger.Errorf("[getProblemAndSubmitId Exists] err:%v", err)
			return nil, fmt.Errorf("internal err")
		}

		res, err := g.DbSubmit.SMembers(g.RedisContext, uidString).Result()
		if err != nil {
			g.Logger.Errorf("[getProblemAndSubmitId SMembers] err:%v", err)
			return nil, fmt.Errorf("internal err")
		}
		return res[1:], nil
	} else {
		//加入默认值，防止脏读
		if _, err = g.DbSubmit.SAdd(g.RedisContext, uidString, define.DefaultRedisValue).Result(); err != nil {
			g.DbSubmit.Del(g.RedisContext, uidString)
			g.Logger.Errorf("[getProblemAndSubmitId SAdd] err:%v", err)
			return nil, fmt.Errorf("internal err")
		}
		if _, err = g.DbSubmit.Expire(g.RedisContext, uidString, define.ExpireTime).Result(); err != nil {
			g.DbSubmit.Del(g.RedisContext, uidString)
			g.Logger.Errorf("[getProblemAndSubmitId Expire] err:%v", err)
			return nil, fmt.Errorf("internal err")
		}
		//不存在，去数据库查询，加入缓存
		submits, err := s.getProblemAndSubmitIdBySql(uid, page, size)
		if err != nil {
			return nil, err
		}
		var res []string
		for _, submit := range submits {
			strSubmitId := strconv.Itoa(submit.Id)
			strProblemId := strconv.Itoa(submit.ProblemId)
			sb := strings.Builder{}
			sb.WriteString(strSubmitId)
			sb.WriteString(" ")
			sb.WriteString(strProblemId)
			res = append(res, sb.String())
			_, err = g.DbSubmit.SAdd(g.RedisContext, uidString, sb.String()).Result()
			if err != nil {
				g.DbSubmit.Del(g.RedisContext, uidString)
				g.Logger.Errorf("[getProblemAndSubmitId SAdd] err:%v", err)
				return nil, fmt.Errorf("internal err")
			}
		}
		return res, nil
	}
}

func (s *sSubmit) getProblemAndSubmitIdBySql(uid, page, size int) ([]*model.Submit, error) {
	//goland:noinspection SqlResolve
	sqlStr := "select id,problem_id from submit where user_id = ? limit ?,?"

	stmt, err := g.MysqlDB.Prepare(sqlStr)
	defer stmt.Close()
	if err != nil {
		g.Logger.Errorf("[getProblemAndSubmitIdBySql] prepare failed,err:%v", err)
		return nil, fmt.Errorf("internal err")
	}
	var submits []*model.Submit
	rows, err := stmt.Query(uid, page, size)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, fmt.Errorf("submit not exist")
		}
		g.Logger.Errorf("[getProblemAndSubmitIdBySql] query failed,err:%v", err)
		return nil, fmt.Errorf("internal err")
	}
	defer rows.Close()
	for rows.Next() {
		var submit model.Submit
		err = rows.Scan(&submit.Id, &submit.ProblemId)
		if err != nil {
			g.Logger.Errorf("[getProblemAndSubmitIdBySql] scan failed,err:%v", err)
			return nil, fmt.Errorf("internal err")
		}
		submits = append(submits, &submit)
	}
	return submits, nil
}

func (s *sSubmit) GetSubmitById(id int) (*model.Submit, error) {
	//goland:noinspection SqlResolve
	sqlStr := "select id,identity,problem_id,user_id,path,status,created_at from submit where id = ?"

	stmt, err := g.MysqlDB.Prepare(sqlStr)
	defer stmt.Close()
	if err != nil {
		g.Logger.Errorf("[getProblemAndSubmitIdBySql] prepare failed,err:%v", err)
		return nil, fmt.Errorf("internal err")
	}
	var submit model.Submit
	err = stmt.QueryRow(id).Scan(&submit.Id, &submit.Identity, &submit.ProblemId, &submit.UserId, &submit.Path, &submit.Status, &submit.CreatedAt)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, fmt.Errorf("sbumit not exitst")
		}
		g.Logger.Errorf("[getProblemAndSubmitIdBySql] query failed,err:%v", err)
		return nil, fmt.Errorf("internal err")
	}
	return &submit, nil
}

func (s *sSubmit) CreatSubmit(submit *model.Submit) error {
	//goland:noinspection SqlResolve
	sqlStr := "insert into submit(identity,problem_id,user_id,path,status,created_at) values(?,?,?,?,?,?)"
	stmt, err := g.MysqlDB.Prepare(sqlStr)
	defer stmt.Close()
	if err != nil {
		g.Logger.Errorf("[CreatSubmit] prepare failed,err:%v", err)
		return fmt.Errorf("internal err")
	}
	_, err = stmt.Exec(submit.Identity, submit.ProblemId, submit.UserId, submit.Path, submit.Status, submit.CreatedAt)
	if err != nil {
		g.Logger.Errorf("[CreatSubmit] update failed,err:%v", err)
		return fmt.Errorf("internal err")
	}
	return nil
}

func (s *sSubmit) CheckCode(testCase *model.TestCase, problem *model.Problem, path string, check int) (int, string) {
	//答案错误的channel
	WA := make(chan int)
	//超内存的channel
	OOM := make(chan int)
	//编译错误的channel
	CE := make(chan int)
	//通过的个数
	var passCount int64
	passCount = 0
	msg := ""
	var status int
	go func() {
		cmd := exec.Command("go", "run", path)
		var out, stderr bytes.Buffer
		cmd.Stderr = &stderr
		cmd.Stdout = &out
		stdinPipe, err := cmd.StdinPipe()
		if err != nil {
			g.Logger.Errorf("StdinPipe err:%v", err)
			return
		}
		io.WriteString(stdinPipe, testCase.Input)
		var bm runtime.MemStats
		runtime.ReadMemStats(&bm)
		//编译出错
		if err := cmd.Run(); err != nil {
			g.Logger.Errorf("run failed,err:%v", stderr.String())
			if err.Error() == "exit status 2" {
				msg = stderr.String()
				CE <- 1
				return
			}
		}
		var em runtime.MemStats
		runtime.ReadMemStats(&em)
		//答案错误
		if testCase.Output != out.String() {
			msg = "答案错误"
			WA <- 1
			return
		}
		//运行超过内存
		if em.Alloc/1024-(bm.Alloc/1024) > uint64(problem.MaxMem) {
			msg = "运行超过内存"
			OOM <- 1
			return
		}

		atomic.AddInt64(&passCount, 1)
	}()

	select {
	case <-WA:
		status = 2
	case <-OOM:
		status = 4
	case <-time.After(time.Millisecond * time.Duration(problem.MaxRuntime)):
		if passCount == int64(check) {
			status = 1
		} else {
			msg = "运行超时"
			status = 3
		}
	}
	return status, msg
}
