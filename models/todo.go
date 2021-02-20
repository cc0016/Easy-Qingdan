package models

import (
	"bubbleList/dao"
)

//清单列表
type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}
//增删改查
func Createdata(todo *Todo)(err error) {
	if err = dao.Db.Create(&todo).Error; err!=nil {
		return err
	}
	return
}

func Getdata(todoList *[]Todo)(err error) {
	if err := dao.Db.Find(&todoList).Error; err!=nil{
		return err
	}
	return
}

func Finddata(id *string, todo *Todo)(err error) {
	if err := dao.Db.Where("id=?", id).First(&todo).Error; err!=nil{
		return err
	}
	return
}

func Putdata(todo *Todo)(err error){
	if err := dao.Db.Save(&todo).Error; err!=nil{
		return err
	}
	return
}

func Deledata(id *string)(err error) {
	if err := dao.Db.Where("id=?",id).Delete(Todo{}).Error; err!=nil{
		return err
	}
	return
}

