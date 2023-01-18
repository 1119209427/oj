package route

import (
	"github.com/gin-gonic/gin"
	g "oj/app/global"
	"oj/app/internal/middlerware"
)

func InitRoute() *gin.Engine {
	r := gin.Default()
	r.Use(middlerware.ZapLogger(g.Logger), middlerware.ZapRecovery(g.Logger, true))
	r.Use(middlerware.Cors())

	routeGroup := new(Group)

	PublishGroup := r.Group("/api")
	{

		routeGroup.InitGetProblemsRoute(PublishGroup)
		routeGroup.InitUserSignRoute(PublishGroup)
	}
	PrivateGroup := r.Group("/api")
	PrivateGroup.Use(middlerware.Auth())
	{
		routeGroup.InitSubmitRoute(PrivateGroup)
		routeGroup.InitProblemRoute(PrivateGroup)
		routeGroup.InitCategoryRoute(PrivateGroup)
		routeGroup.InitTestCaeRoute(PrivateGroup)
	}
	InternalGroup := r.Group("/api")
	{
		routeGroup.InitUserAdminRoute(InternalGroup)

	}

	g.Logger.Infof("initialize routers successfully")
	return r

}
