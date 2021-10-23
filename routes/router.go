package routes

import (
	v1 "ginBlog/api/v1"
	"ginBlog/middleware"
	"ginBlog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter()  {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())

	{
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		auth.POST("category/add",v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)

		auth.POST("article/add",v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
	}

	routerV1 := r.Group("api/v1")
	{
		routerV1.POST("user/add",v1.AddUser)
		routerV1.GET("users", v1.GetUser)
		routerV1.GET("categories", v1.GetCategory)
		routerV1.GET("articles", v1.GetArticle)
		routerV1.GET("article/list/:id", v1.GetCategoryArticle)
		routerV1.GET("article/info/:id", v1.GetArticleInfo)
		routerV1.POST("login", v1.Login)

	}


	_ = r.Run(utils.HttpPort)
}
