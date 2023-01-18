package route

import (
	"github.com/gin-gonic/gin"
	"oj/app/api"
	"oj/app/internal/middlerware"
)

type CategoryRoute struct {
}

func (r *CategoryRoute) InitCategoryRoute(route *gin.RouterGroup) (R gin.IRoutes) {
	categoryRoute := route.Group("/category")
	categoryApi := api.Category()
	{
		categoryRoute.POST("/created", middlerware.Admin(), categoryApi.Category().CreateCategory)
		categoryRoute.PUT("/update", middlerware.Admin(), categoryApi.Category().UpdateCategory)
	}
	return categoryRoute
}
