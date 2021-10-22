package v1

import (
	"ginBlog/model"
	"ginBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func AddCategory(c *gin.Context) {

	var catogory model.Category
	_ = c.ShouldBindJSON(&catogory)
	code = model.CheckCategory(catogory.Name)
	if code == errmsg.SUCCSE {
		model.CreateCategory(&catogory)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		code = errmsg.ERROR_CATENAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    catogory,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetCategory(c *gin.Context) {
	paseSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if paseSize == 0 {
		paseSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data := model.GetCategories(paseSize, pageNum)
	code = errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func EditCategory(c *gin.Context) {
	var category model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&category)
	code = model.CheckCategory(category.Name)
	if code == errmsg.SUCCSE {
		model.EditCategory(id, &category)
	}

	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteCategory(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

