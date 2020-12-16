package models

import (
	database "lirawx.cn/go-web/db"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/*
	Todo这个Model的增删改查操作都放在这里
*/
// CreateATodo 创建todo
func CreateATodo(todo *Todo) (err error) {
	err = database.DB.Create(&todo).Error
	return
}

// GetAllTodo 获取todo
func GetAllTodo() (todoList []*Todo, err error) {
	if err = database.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

// GetATodo 获取一个todo
func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = database.DB.Debug().Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

// UpdateATodo 更新一个todo
func UpdateATodo(todo *Todo) (err error) {
	err = database.DB.Save(todo).Error
	return
}

// DeleteATodo 删除一个todo
func DeleteATodo(id string) (err error) {
	err = database.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
