package v1

import (
	"ginBlog/model"
	"ginBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

func AddUser(c *gin.Context)  {

	var data model.User
	_  = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data" : data,
		"message" : errmsg.GetErrMsg(code),
	})
}

func GetUser(c *gin.Context)  {
	paseSize, _:= strconv.Atoi(c.Query("pagesize"))
	pageNum,_ := strconv.Atoi(c.Query("pagenum"))

	if paseSize == 0 {
		paseSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data := model.GetUsers(paseSize,pageNum)
	code = errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"status":code,
		"data": data,
		"message":errmsg.GetErrMsg(code),
	})
}

func EditUser(c *gin.Context)  {
	
}

func DeleteUser	(c *gin.Context)  {
	
}