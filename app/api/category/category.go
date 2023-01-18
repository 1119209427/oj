package category

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oj/app/define"
	g "oj/app/global"
	"oj/app/internal/model"
	"oj/app/internal/service"
	"strconv"
	"time"
)

type CateApi struct{}

var insCate = CateApi{}

type ResList struct {
	category       *model.Category
	categoryParent *model.Category
}

func (a *CateApi) GetCategoryList(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		g.Logger.Errorf("GetCategoryList Page strconv Error %v", err)
		return
	}
	page = (page - 1) * size
	keyword := c.Query("keyword")
	categories, err := service.Category().CategoryService().GetCategoryLists(page, size, keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
			"ok":   false,
		})
		return
	}
	var lists []*ResList
	for _, category := range categories {
		var list *ResList
		if category.Id != define.ParentDefault {
			categoryParent, err := service.Category().CategoryService().GetCategoryInfoById(category.ParentId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  err.Error(),
					"ok":   false,
				})
				return
			}
			list.category = category
			list.categoryParent = categoryParent
			lists = append(lists, list)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  lists,
		"ok":   true,
	})
}

func (a *CateApi) UpdateCategory(c *gin.Context) {
	idStr := c.PostForm("id")
	name := c.PostForm("name")
	parentIdStr := c.DefaultPostForm("parent_id", define.ParentId)
	id, _ := strconv.Atoi(idStr)
	parentId, err := strconv.Atoi(parentIdStr)
	if err != nil {
		g.Logger.Errorf("ShouldBind err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "internal err",
			"ok":   false,
		})
		return
	}
	var category model.Category
	category.Id = id
	category.Name = name
	category.ParentId = parentId
	err = service.Category().CategoryService().UpdateCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
			"ok":   false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "update category success",
		"ok":   true,
	})
}

func (a *CateApi) CreateCategory(c *gin.Context) {
	name := c.PostForm("name")
	parentId := c.DefaultPostForm("parent_id", define.ParentId)
	id, err := strconv.Atoi(parentId)
	if err != nil {
		g.Logger.Errorf("ShouldBind err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "internal err",
			"ok":   false,
		})
		return
	}
	var category model.Category
	category.Name = name
	category.ParentId = id
	category.CreatedAt = time.Now()
	err = service.Category().CategoryService().CreateCategory(category)
	if err != nil {
		g.Logger.Errorf("ShouldBind err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
			"ok":   false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "create category success",
		"ok":   true,
	})
}
