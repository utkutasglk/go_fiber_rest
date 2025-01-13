package main

import (
	"go_fiber/dal"
	"go_fiber/database"

	"github.com/gofiber/fiber/v2"
)

func main(){

	database.Connect()
	database.DB.AutoMigrate(dal.Todo{})

	app := fiber.New()
 
	app.Get("/", func(c *fiber.Ctx) error{

		return c.SendString("Hello World")
	})


	app.Listen("localhost:3000")
	
}