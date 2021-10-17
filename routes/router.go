package routes

import (
	v1 "ginBlog/api/v1"
	"ginBlog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter()  {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	routerV1 := r.Group("api/v1")
	{
		routerV1.POST("user/add",v1.AddUser)
		routerV1.GET("users", v1.GetUser)
		routerV1.PUT("user/:id", v1.EditUser)
		routerV1.DELETE("user/:id", v1.DeleteUser)
	}

	r.Run(utils.HttpPort)
}
