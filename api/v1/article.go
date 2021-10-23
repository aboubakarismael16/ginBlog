package v1

import (
	"ginBlog/model"
	"ginBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddArticle(c *gin.Context) {

	var article model.Article
	_ = c.ShouldBindJSON(&article)

	code = model.CreateArticle(&article)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    article,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetCategoryArticle(c *gin.Context)  {
	paseSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))

	if paseSize == 0 {
		paseSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, code := model.GetCategoryArticle(id,paseSize,pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetArticleInfo(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArticleInfo(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetArticle(c *gin.Context) {
	paseSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if paseSize == 0 {
		paseSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data ,code := model.GetArticle(paseSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func EditArticle(c *gin.Context) {
	var article model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&article)

	code = model.EditArticle(id, &article)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteArticle(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

