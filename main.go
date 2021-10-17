package main

import (
	"ginBlog/model"
	"ginBlog/routes"
)

func main() {
	model.InitDb()

	routes.InitRouter()
}
