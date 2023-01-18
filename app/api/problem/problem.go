package problem

import (
	"github.com/gin-gonic/gin"
	"net/http"
	g "oj/app/global"
	"oj/app/internal/model"
	"oj/app/internal/service"
	"time"
)

type BasicProblem struct {
	Title      string `form:"title" json:"title" binding:"required"`
	Content    string `form:"content" json:"content" binding:"required"`
	TextCase   string `form:"textCase" json:"textCase" binding:"required"`
	MaxRuntime int    `form:"max_runtime" json:"max_runtime" binding:"required"`
	MaxMem     int    `form:"max_mem" json:"max_mem" binding:"required"`
}

type ApiProblem struct {
}

var insProblem = ApiProblem{}

func (a *ApiProblem) CreateProblem(c *gin.Context) {
	var problemBasic BasicProblem
	err := c.ShouldBind(&problemBasic)
	if err != nil {
		g.Logger.Errorf("ShouldBind err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "internal err",
			"ok":   false,
		})
		return
	}
	var problem model.Problem
	problem.Identity = service.Problem().ProblemService().UUid()
	problem.Title = problemBasic.Title
	problem.Content = problemBasic.Content
	problem.TextCase = problemBasic.TextCase
	problem.MaxRuntime = problemBasic.MaxRuntime
	problem.MaxMem = problemBasic.MaxMem
	problem.CreatedAt = time.Now()

	err = service.Problem().ProblemService().CreateProblem(&problem)
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
		"msg":  "create problem success",
		"ok":   true,
	})
}

type ChangeProblem struct {
	Id         int    `form:"id" json:"id" binding:"required"`
	Title      string `form:"title" json:"title" binding:"required"`
	Content    string `form:"content" json:"content" binding:"required"`
	TextCase   string `form:"textCase" json:"textCase" binding:"required"`
	MaxRuntime int    `form:"max_runtime" json:"max_runtime" binding:"required"`
	MaxMem     int    `form:"max_mem" json:"max_mem" binding:"required"`
}

func (a *ApiProblem) ChangeProblem(c *gin.Context) {
	var changeProblem ChangeProblem
	err := c.ShouldBind(&changeProblem)
	if err != nil {
		g.Logger.Errorf("ShouldBind err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "internal err",
			"ok":   false,
		})
		return
	}
	var problem model.Problem
	problem.Id = changeProblem.Id
	problem.Title = changeProblem.Title
	problem.Content = changeProblem.Content
	problem.TextCase = changeProblem.TextCase
	problem.MaxRuntime = changeProblem.MaxRuntime
	problem.MaxMem = changeProblem.MaxMem
	err = service.Problem().ProblemService().ChangeProblem(&problem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "internal err",
			"ok":   false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "update problem success",
		"ok":   true,
	})
}
