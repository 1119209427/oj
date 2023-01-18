package route

import (
	"github.com/gin-gonic/gin"
	"oj/app/api"
	"oj/app/internal/middlerware"
)

type TestCaseRoute struct {
}

func (r *TestCaseRoute) InitTestCaeRoute(route *gin.RouterGroup) (R gin.IRoutes) {
	testCaseRoute := route.Group("/testcase")
	testCaseApi := api.Problem().TestCase()
	{
		testCaseRoute.POST("/create", middlerware.Admin(), testCaseApi.CreateTestCase)
		testCaseRoute.PUT("/update", middlerware.Admin(), testCaseApi.UpdateTestCase)
	}
	return testCaseRoute
}
