package problem

import (
	"github.com/gin-gonic/gin"
	"net/http"
	g "oj/app/global"
	"oj/app/internal/model"
	"oj/app/internal/service"
)

type TestApi struct {
}

var insTestCase = TestApi{}

type TestCaseBasic struct {
	ProblemId int    `form:"problem_id" json:"problem_id" binding:"required"`
	Input     string `form:"input" json:"input" binding:"required"`
	Output    string `form:"output" json:"output" binding:"required"`
}

func (a *TestApi) CreateTestCase(c *gin.Context) {
	var testBasic TestCaseBasic
	err := c.ShouldBind(&testBasic)
	if err != nil {
		g.Logger.Errorf("ShouldBind err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "internal err",
			"ok":   false,
		})
		return
	}
	var testCase model.TestCase
	testCase.ProblemId = testBasic.ProblemId
	testCase.Input = testBasic.Input
	testCase.Output = testBasic.Output
	err = service.TestCase().TestCaseService().CreateTestCase(&testCase)
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
		"msg":  "create testcase success",
		"ok":   true,
	})
}

type TestCaseChange struct {
	Id        int    `form:"id" json:"id" binding:"required"`
	ProblemId int    `form:"problem_id" json:"problem_id" binding:"required"`
	Input     string `form:"input" json:"input" binding:"required"`
	Output    string `form:"output" json:"output" binding:"required"`
}

func (a *TestApi) UpdateTestCase(c *gin.Context) {
	var testBasic TestCaseChange
	err := c.ShouldBind(&testBasic)
	if err != nil {
		g.Logger.Errorf("ShouldBind err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "internal err",
			"ok":   false,
		})
		return
	}
	var testCase model.TestCase
	testCase.Id = testBasic.Id
	testCase.ProblemId = testBasic.ProblemId
	testCase.Input = testBasic.Input
	testCase.Output = testBasic.Output
	err = service.TestCase().TestCaseService().UpdateTestCase(&testCase)
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
		"msg":  "update testcase success",
		"ok":   true,
	})
}
