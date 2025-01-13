package services

import (
	"errors"
	"fmt"
	"go_fiber/dal"
	"go_fiber/types"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var validate = validator.New()


func CreateTodo(c *fiber.Ctx) error{

		t := new(types.TodoCreateDTO)

		if err := c.BodyParser(t);  err != nil{
			return c.Status(400).JSON(fiber.Map{
				"message":"Bad Request",
			})
		}	
   
		if err := validate.Struct(t); err != nil{
     
			valErrors := err.(validator.ValidationErrors)[0]
			message := fmt.Sprintf("Field: '%s', failed on '%s' with your value: '%s'", valErrors.Field(), valErrors.Tag(), valErrors.Value() )
			return c.Status(400).JSON(fiber.Map{
				"message":message,
			})
		}

		newTodo := dal.Todo{
			Title: t.Title,
		}

		if res := dal.CreateTodo(&newTodo); res.Error != nil{
			return res.Error
		}

		return c.JSON(fiber.Map{
			"message":"Todo created successfully",
		})


	}

	func GetTodos(c *fiber.Ctx) error{

		todos := []types.TodoResponse{}
		
		res := dal.GetTodos(&todos)
		
		if res.Error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message":"Bad Request",
		})

		}

		return c.JSON(todos)
	}

	func GetTodoById(c *fiber.Ctx) error{

		todoID := c.Params("todoID")

		d := types.TodoResponse{}

		res := dal.GetTodoById(&d, todoID)
		if res.Error != nil {
			if errors.Is(res.Error, gorm.ErrRecordNotFound) {

				return c.Status(404).JSON(fiber.Map{
				"message":"Todo not found",
		     })
			
		}
				return c.Status(500).JSON(fiber.Map{
				"message":"Failed to get todo",
		     })
	}
	return c.JSON(d)
}

func UpdateTodo(c *fiber.Ctx) error{

			todoID := c.Params("todoID")

			t := new(types.TodoUpdateDTO)

			if err := c.BodyParser(t);  err != nil{
			return c.Status(400).JSON(fiber.Map{
				"message":"Bad Request",
			})
		}	
   
		if err := validate.Struct(t); err != nil{
     
			valErrors := err.(validator.ValidationErrors)[0]
			message := fmt.Sprintf("Field: '%s', failed on '%s' with your value: '%s'", valErrors.Field(), valErrors.Tag(), valErrors.Value() )
			return c.Status(400).JSON(fiber.Map{
				"message":message,
			})
		}

		res := dal.UpdateTodo(todoID, t)

		if res.Error != nil || res.RowsAffected == 0{

			return c.Status(500).JSON(fiber.Map{
				"message":"Failed to update todo",
			})
		}

		return c.JSON(fiber.Map{
			"message":"Todo Updated Successfully",
		})



}