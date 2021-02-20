package main

import (
	"bubbleList/dao"
	"bubbleList/models"
	"bubbleList/router"
	"fmt"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	err := dao.Init()
	if err != nil {
		fmt.Println("err: ", err)
		panic(err)
	}
	defer dao.Db.Close()
	dao.Db.AutoMigrate(&models.Todo{})

	r := router.SetupRouter()

	r.Run(":9090")
}
