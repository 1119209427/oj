package middlerware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthBody() gin.HandlerFunc {
	return func(context *gin.Context) {
		auth := context.Request.PostFormValue("token")
		fmt.Printf("%v \n", auth)

		if len(auth) == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "not logged in",
				"ok":   false,
			})
		}
		auth = strings.Fields(auth)[1]
		token, err := parseToken(auth)
		if err != nil {
			context.Abort()
			context.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
				"ok":   false,
			})
		} else {
			println("token 正确")
		}
		context.Set("userId", token.Id)
		context.Next()
	}
}
