package route

import (
	"github.com/gin-gonic/gin"
	"oj/app/api"
)

type UserRoute struct{}

func (r *UserRoute) InitUserSignRoute(route *gin.RouterGroup) (R gin.IRoutes) {
	userRoute := route.Group("/user")
	userApi := api.User()
	{
		userRoute.POST("/register", userApi.User().Register)
		userRoute.POST("/login", userApi.User().Login)
	}
	return userRoute
}
func (r *UserRoute) InitUserAdminRoute(route *gin.RouterGroup) (R gin.IRoutes) {
	userRoute := route.Group("/user")
	userApi := api.User()
	{
		userRoute.PUT("/admin", userApi.Admin().AddAdmin)
		userRoute.PUT("/cancel admin", userApi.Admin().CancelAdmin)
	}
	return userRoute
}
