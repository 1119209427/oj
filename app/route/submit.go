package route

import (
	"github.com/gin-gonic/gin"
	"oj/app/api"
)

type SubmitRoute struct{}

func (r *SubmitRoute) InitSubmitRoute(route *gin.RouterGroup) (R gin.IRoutes) {
	submitRoute := route.Group("/submit")
	submitApi := api.Submit()
	{
		submitRoute.GET("/lists", submitApi.Submit().GetSubmitList)
		submitRoute.POST("/submit", submitApi.Submit().Submit)
	}
	return submitRoute
}
