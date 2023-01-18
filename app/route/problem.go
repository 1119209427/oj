package route

import (
	"github.com/gin-gonic/gin"
	"oj/app/api"
	"oj/app/internal/middlerware"
)

type ProblemRoute struct{}

func (r *ProblemRoute) InitGetProblemsRoute(route *gin.RouterGroup) (R gin.IRoutes) {
	problemRoute := route.Group("/problem")
	problemApi := api.Problem()
	{
		problemRoute.GET("/list", problemApi.ProblemList().GetProblemList)
	}

	return problemRoute
}

func (r *ProblemRoute) InitProblemRoute(route *gin.RouterGroup) (R gin.IRoutes) {
	problemRoute := route.Group("/problem")
	problemApi := api.Problem()
	{
		problemRoute.POST("/create", middlerware.Admin(), problemApi.Problem().CreateProblem)
		problemRoute.PUT("/update", middlerware.Admin(), problemApi.Problem().ChangeProblem)
	}
	return problemRoute
}
