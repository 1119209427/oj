package submit

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"oj/app/define"
	g "oj/app/global"
	"oj/app/internal/model"
	"oj/app/internal/service"
	"oj/utils/file"
	"oj/utils/uuid"
	"strconv"
	"strings"
	"time"
)

type Api struct{}

var insSubmit = Api{}

type Response struct {
	Lists []*List
}
type List struct {
	Problem *model.Problem
	Submit  *model.Submit
}

func (a *Api) GetSubmitList(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		g.Logger.Errorf("GetProblemList Page strconv Error %v", err)
		return
	}
	page = (page - 1) * size
	uidStr := c.GetString("userId")
	uid, err := strconv.Atoi(uidStr)
	var res Response
	if err != nil {
		g.Logger.Errorf("GetSubmitList Page strconv Error %v", err)
		return
	}
	ids, err := service.Submit().Submit().GetProblemAndSubmitIdByUserId(uid, page, size)
	if err != nil {
		switch err.Error() {
		case "internal err":
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})
		case "submit not exist":
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  err.Error(),
				"ok":   true,
			})
		}
		return
	}
	var lists []*List
	for _, id := range ids {
		var list List
		str := strings.Split(id, " ")
		strSubmitId := str[0]
		strProblemId := str[1]
		submitId, _ := strconv.Atoi(strSubmitId)
		problemId, err := strconv.Atoi(strProblemId)
		if err != nil {
			g.Logger.Errorf("GetProblemList Page strconv Error %v", err)
			return
		}
		submit, err := service.Submit().Submit().GetSubmitById(submitId)
		if err != nil {
			switch err.Error() {
			case "internal err":
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  err.Error(),
					"ok":   false,
				})
			case "submit not exist":
				c.JSON(http.StatusOK, gin.H{
					"code": http.StatusOK,
					"msg":  err.Error(),
					"ok":   true,
				})
			}
			return
		}

		problem, err := service.Problem().ProblemService().GetProblemById(problemId)
		if err != nil {
			switch err.Error() {
			case "internal err":
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  err.Error(),
					"ok":   false,
				})
			case "problem not exist":
				c.JSON(http.StatusOK, gin.H{
					"code": http.StatusOK,
					"msg":  err.Error(),
					"ok":   true,
				})
			}
			return
		}

		list.Submit = submit
		list.Problem = problem
		lists = append(lists, &list)
	}
	res.Lists = lists
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  res,
		"ok":   true,
	})

}

func (a *Api) Submit(c *gin.Context) {
	strId := c.Query("problem_id")
	uid, _ := strconv.Atoi(c.GetString("userId"))
	problemId, err := strconv.Atoi(strId)
	if err != nil {
		g.Logger.Errorf("strconv err:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "internal err",
			"ok":   false,
		})
		return
	}
	code, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		g.Logger.Errorf("ReadAll err:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "read code failed",
			"ok":   false,
		})
		return
	}
	problem, err := service.Problem().ProblemService().GetProblemById(problemId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
			"ok":   false,
		})
		return
	}
	//保存代码到本地
	path, err := file.CodeSave(code)
	//检查代码
	//step1:获取到测试案例
	testCases, err := service.TestCase().TestCaseService().GetTestCase(problemId)
	check := len(testCases)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
			"ok":   false,
		})
		return
	}
	//step2:判读代码
	var status int
	var msg string
	for _, testCase := range testCases {
		status, msg = service.Submit().Submit().CheckCode(testCase, problem, path, check)
		if msg != "" {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  msg,
				"ok":   false,
			})
			return
		}
	}
	//step3:存储到数据库
	var submit model.Submit
	submit.Status = status
	submit.Path = path
	submit.ProblemId = problemId
	submit.UserId = uid
	submit.Identity = uuid.GetUUid()
	submit.CreatedAt = time.Now()
	err = service.Submit().Submit().CreatSubmit(&submit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
			"ok":   false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "代码提交正确",
		"ok":   true,
	})
}
