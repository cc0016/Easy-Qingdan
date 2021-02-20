package router

import (
	"bubbleList/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "static")

	r.GET("/", controller.Indexhandler)

	v1Group := r.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", controller.CreateAtodo)
		//查看所有待办事项
		v1Group.GET("/todo", controller.GetList)
		//修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.PutList)
		//删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleList)
	}
	return r
}
