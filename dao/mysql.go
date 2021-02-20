package dao

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Db *gorm.DB
)

func Init() (err error){
	dsn := "root:chenyuhan123000@(localhost:3306)/bubble?parseTime=True&loc=Local"
	Db, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return Db.DB().Ping()
}
