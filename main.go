package main

import (
	"go_fiber/dal"
	"go_fiber/database"
	"go_fiber/services"

	"github.com/gofiber/fiber/v2"
)


func main(){

	database.Connect()
	database.DB.AutoMigrate(&dal.Todo{})

	app := fiber.New()
 
	app.Get("/", func(c *fiber.Ctx) error{
		return c.JSON(fiber.Map{
			"message":"Hello world",
		})
	})

	app.Post("/todos",services.CreateTodo)

	app.Get("/todos", services.GetTodos)
 
	app.Listen("localhost:3000")
	
}