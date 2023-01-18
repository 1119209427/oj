package middlerware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	g "oj/app/global"
	"oj/app/internal/service"
	"strconv"
)

func Admin() gin.HandlerFunc {
	return func(context *gin.Context) {
		strId := context.GetString("userId")
		id, err := strconv.Atoi(strId)
		if err != nil {
			g.Logger.Errorf("parase err:%v", err)
			context.Abort()
			context.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "internal err",
				"ok":   false,
			})
			return
		}
		err = service.User().User().CheckAdmin(id)
		if err != nil {
			context.Abort()
			switch err.Error() {
			case "internal err":
				context.JSON(http.StatusInternalServerError, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  err.Error(),
					"ok":   false,
				})
			case "not have this root":
				context.JSON(http.StatusUnauthorized, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  err.Error(),
					"ok":   false,
				})
			}
			return
		}
		context.Next()
	}
}
