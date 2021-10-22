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

		routerV1.POST("category/add",v1.AddCategory)
		routerV1.GET("categories", v1.GetCategory)
		routerV1.PUT("category/:id", v1.EditCategory)
		routerV1.DELETE("category/:id", v1.DeleteCategory)
	}

	r.Run(utils.HttpPort)
}
