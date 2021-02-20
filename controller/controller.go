package controller

import (
	"bubbleList/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Indexhandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html",nil)
}

func CreateAtodo(c *gin.Context) {
	//取前端json数据并在数据库中创建数据
	var todo models.Todo
	c.BindJSON(&todo)

	err := models.Createdata(&todo)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}else{
		c.JSON(http.StatusOK, todo)
	}
}

func GetList(c *gin.Context) {
	//查询todo表中的所有数据并返回
	var todoList []models.Todo

	err := models.Getdata(&todoList)
	if err!=nil{
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}else {
		c.JSON(http.StatusOK, todoList)
	}
}

func PutList(c *gin.Context) {
	//拿到id并查找数据库进行修改
	id, ok := c.Params.Get("id")
	if !ok{
		c.JSON(http.StatusOK, gin.H{
			"error": "Invalid id",
		})
		return
	}
	var todo models.Todo
	err := models.Finddata(&id, &todo)
	if err!=nil{  //查询数据库
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}
	c.BindJSON(&todo)
	err = models.Putdata(&todo)
	if err!=nil{  //更新数据库
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}else {
		c.JSON(http.StatusOK, todo)
	}//更新数据库
}

func DeleList(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok{
		c.JSON(http.StatusOK, gin.H{
			"error": "Invalid id",
		})
		return
	}

	err := models.Deledata(&id)
	if err!=nil{
		c.JSON(http.StatusOK, gin.H{"error": "Invalid id"})
	} else {
		c.JSON(http.StatusOK, gin.H{id:"deleted"})
	}
}