package router

import (
	"memo/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() (*gin.Engine) {
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		// 添加事项
		v1Group.POST("/todo", controller.CreateATodo)
		// 查看所有代办事项
		v1Group.GET("/todo", controller.GetTodoList)
		//修改某一个代办事项
		v1Group.PUT("/todo/:id", controller.UpdateTodoStatus)
		//删除某一个代办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}

	return r
}