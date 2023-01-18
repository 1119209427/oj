package problem

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	g "oj/app/global"
	"oj/app/internal/model"
	"sync"
)

type sProblem struct{}

var (
	onceProblem sync.Once
	insProblem  *sProblem
)

func newProblemService() *sProblem {
	onceProblem.Do(
		func() {
			insProblem = &sProblem{}
		})
	return insProblem
}

func (s *sProblem) ChangeProblem(problem *model.Problem) error {
	//goland:noinspection SqlResolve
	sqlStr := "update problem set title=?,content=?,test_case=?,max_runtime=?,max_mem=? where id=?"
	stmt, err := g.MysqlDB.Prepare(sqlStr)
	if err != nil {
		g.Logger.Errorf("[ChangeProblem] prepare failed,err:%v", err)
		return fmt.Errorf("internal err")
	}
	defer stmt.Close()
	_, err = stmt.Exec(problem.Title, problem.Content, problem.TextCase, problem.MaxRuntime, problem.MaxMem, problem.Id)
	if err != nil {
		g.Logger.Errorf("[ChangeProblem] update failed,err:%v", err)
		return fmt.Errorf("internal err")
	}
	return nil
}

func (s *sProblem) CreateProblem(problem *model.Problem) error {
	//goland:noinspection SqlResolve
	sqlStr := "insert into problem(identity,title,content,test_case,max_runtime,max_mem,created_at) values(?,?,?,?,?,?,?)"
	stmt, err := g.MysqlDB.Prepare(sqlStr)
	if err != nil {
		g.Logger.Errorf("[CreateProblem] prepare failed,err:%v", err)
		return fmt.Errorf("internal err")
	}
	defer stmt.Close()
	_, err = stmt.Exec(problem.Identity, problem.Title, problem.Content, problem.TextCase, problem.MaxRuntime, problem.MaxMem, problem.CreatedAt)
	if err != nil {
		g.Logger.Errorf("[CreateProblem] insert failed,err:%v", err)
		return fmt.Errorf("internal err")
	}
	return nil
}

func (s *sProblem) GetProblemList(ctx context.Context, keyword string, size, page int) ([]*model.Problem, int, error) {
	var problems []*model.Problem
	//goland:noinspection SqlResolve
	sqlStr := "select id,identity,max_runtime,max_mem,title,created_at from problem where (title like ? or content like ?) limit ?,?"
	stmt, err := g.MysqlDB.Prepare(sqlStr)
	if err != nil {
		g.Logger.Errorf("[GetProblemList] prepare failed,err:%v", err)
		return nil, 0, fmt.Errorf("internal err")
	}
	defer stmt.Close()
	keyword = "%" + keyword + "%"
	rows, err := stmt.Query(keyword, keyword, page, size)
	if err != nil {
		g.Logger.Errorf("[GetProblemList] query failed,err:%v", err)
		return nil, 0, fmt.Errorf("internal err")
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		var problem model.Problem
		err = rows.Scan(&problem.Id, &problem.Identity, &problem.MaxRuntime, &problem.MaxMem, &problem.Title, &problem.CreatedAt)
		if err != nil {
			g.Logger.Errorf("[GetProblemList] scan failed,err:%v", err)
			return nil, 0, fmt.Errorf("internal err")
		}
		count++
		problems = append(problems, &problem)
	}
	return problems, count, nil
}

func (s *sProblem) GetProblemById(id int) (*model.Problem, error) {
	//goland:noinspection SqlResolve
	sqlStr := "select id,identity,max_runtime,max_mem,title,created_at from problem where id = ?"
	stmt, err := g.MysqlDB.Prepare(sqlStr)
	if err != nil {
		g.Logger.Errorf("[GetProblemById] prepare failed,err:%v", err)
		return nil, fmt.Errorf("internal err")
	}
	defer stmt.Close()
	var problem model.Problem
	err = stmt.QueryRow(id).Scan(&problem.Id, &problem.Identity, &problem.MaxRuntime, &problem.MaxMem, &problem.Title, &problem.CreatedAt)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, fmt.Errorf("problem not exist")
		}
		g.Logger.Errorf("[GetProblemById] prepare failed,err:%v", err)
		return nil, fmt.Errorf("internal err")
	}
	return &problem, nil
}

func (s *sProblem) UUid() string {
	return uuid.NewV4().String()
}
