package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	g "oj/app/global"
	"oj/app/internal/model"
	"oj/app/internal/service"
	"strconv"
	"time"
)

type SignApi struct{}

var insSign = SignApi{}

func (a *SignApi) Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "username cannot be null",
			"ok":   false,
		})
		return
	}
	if password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "password cannot be null",
			"ok":   false,
		})
		return
	}

	err := service.User().User().CheckUserIsExist(username)
	if err != nil {
		if err.Error() == "internal err" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})
		} else if err.Error() == "username already exist" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
				"ok":   false,
			})
		}

		return
	}

	user := model.User{}
	salt := strconv.FormatInt(time.Now().Unix(), 10)

	encryptedPassword := service.User().User().EncryptPassword(salt, password)

	user.Name = username
	user.Salt = salt
	user.Password = encryptedPassword
	user.CreatedAt = time.Now()
	user.Identity = service.User().User().UUid()

	err = service.User().User().CreateUser(user)
	if err != nil {
		if err.Error() == "internal err" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "register successfully",
		"ok":   true,
	})
}

func (a *SignApi) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "username cannot be null",
			"ok":   false,
		})
		return
	}
	if password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "password cannot be null",
			"ok":   false,
		})
		return
	}

	userSubject := &model.User{
		Name:     username,
		Password: password,
	}

	err := service.User().User().CheckPassword(userSubject)
	if err != nil {
		switch err.Error() {
		case "internal err":
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})
		case "invalid username or password":
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
				"ok":   false,
			})
		}

		return
	}

	tokenString, err := service.User().User().GenerateToken(c, userSubject)
	g.Logger.Info(tokenString)
	if err != nil {
		switch err.Error() {
		case "internal err":
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "login successfully",
		"token": tokenString,
		"ok":    true,
	})
}
