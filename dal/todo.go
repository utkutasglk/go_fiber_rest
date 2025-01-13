package dal

import (
	"go_fiber/database"

	"gorm.io/gorm"
)

type Todo struct {
	ID int
	Title string
	Completed bool `gorm:"default:false"`
} 

func CreateTodo(todo *Todo) *gorm.DB{
	return database.DB.Create(&todo)
}

func GetTodos(dest any)*gorm.DB{
	return database.DB.Model(&Todo{}).Find(dest)
}