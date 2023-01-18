package problem

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oj/app/define"
	g "oj/app/global"
	"oj/app/internal/model"
	"oj/app/internal/service"
	"strconv"
)

type ListApi struct{}

var insList = ListApi{}

type Response struct {
	ProblemsList *[]*ProblemsList
	Count        int
}
type ProblemsList struct {
	Problem  *model.Problem
	Category []*model.Category
}

func (a *ListApi) GetProblemList(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		g.Logger.Errorf("GetProblemList Page strconv Error %v", err)
		return
	}
	page = (page - 1) * size
	keyword := c.Query("keyword")

	problems, count, err := service.Problem().ProblemService().GetProblemList(c.Request.Context(), keyword, size, page)
	if err != nil {
		switch err.Error() {
		case "internal err":
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})
		}
		return
	}
	var problemsList []*ProblemsList
	for _, problem := range problems {
		var content ProblemsList
		ids, err := service.CategoryProblem().CategoryProblemService().GetCategoryIdByProblemId(problem.Id)
		if err != nil {
			switch err.Error() {
			case "internal err":
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  err.Error(),
					"ok":   false,
				})
			case "category not exist":
				c.JSON(http.StatusBadRequest, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  err.Error(),
					"ok":   false,
				})
			}
			return
		}
		var categories []*model.Category
		for _, id := range ids {
			category, err := service.Category().CategoryService().GetCategoryInfoById(id)
			if err != nil {
				switch err.Error() {
				case "internal err":
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": http.StatusInternalServerError,
						"msg":  err.Error(),
						"ok":   false,
					})
				case "category not exist":
					c.JSON(http.StatusBadRequest, gin.H{
						"code": http.StatusInternalServerError,
						"msg":  err.Error(),
						"ok":   false,
					})
				}
				return
			}
			categories = append(categories, category)
		}
		content.Problem = problem
		content.Category = categories
		problemsList = append(problemsList, &content)
	}
	var res Response
	res.ProblemsList = &problemsList
	res.Count = count

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  res,
		"ok":   true,
	})

}
