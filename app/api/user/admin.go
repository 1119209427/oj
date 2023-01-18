package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	g "oj/app/global"
	"oj/app/internal/service"
	"strconv"
)

type AdminApi struct {
}

var insAdmin = AdminApi{}

func (a *AdminApi) AddAdmin(ctx *gin.Context) {
	//获取到要升级权限的用户
	ids := ctx.PostFormArray("ids")
	for _, sId := range ids {
		id, err := strconv.Atoi(sId)
		if err != nil {
			g.Logger.Errorf("parese err:%v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "传入参数错误",
				"ok":   false,
			})
			return
		}
		err = service.User().User().ChangeUserAdmin(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "change admin successfully",
		"ok":   true,
	})
}

func (a *AdminApi) CancelAdmin(ctx *gin.Context) {
	//获取到取消用户权限的接口
	ids := ctx.PostFormArray("ids")
	for _, sId := range ids {
		id, err := strconv.Atoi(sId)
		if err != nil {
			g.Logger.Errorf("parese err:%v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "传入参数错误",
				"ok":   false,
			})
			return
		}
		err = service.User().User().CancelUserAdmin(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "change admin successfully",
		"ok":   true,
	})
}
