package v1

import (
	"ginBlog/model"
	"ginBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)


func AddUser(c *gin.Context)  {
	// todo
	var data model.User
	_  = c.ShouldBindJSON(&data)
	code := model.CheckUser(data.Username)
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
	
}

func EditUser(c *gin.Context)  {
	
}

func DeleteUser	(c *gin.Context)  {
	
}