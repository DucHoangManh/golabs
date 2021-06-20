package handlers

import (
	"github.com/DucHoangManh/golabs/restAPI/database"
	"github.com/DucHoangManh/golabs/restAPI/models"
	"github.com/gofiber/fiber/v2"
)

func AllCategories(c *fiber.Ctx) error {
	var categories []models.Category
	err := database.DB.Find(&categories).Error
	if err != nil {
		return err
	}
	return c.JSON(categories)
}
func CreateCategory(c *fiber.Ctx) error {
	var category models.Category
	err := c.BodyParser(&category)
	if err != nil {
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message": "can't parse request",
		})
	}
	err = database.DB.Create(&category).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" :err.Error(),
		})
	}
	database.DB.First(&category)
	return c.JSON(fiber.Map{
		"message" : "success",
		"category" : category,
	})
}

func GetCategory (c *fiber.Ctx) error {
	Id, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : "invalid id",
		})
	}
	category:= models.Category{
		Id: uint(Id),
	}
	database.DB.Find(&category)
	return c.JSON(fiber.Map{
		"message" : "success",
		"payload" : fiber.Map{
			"category" : category,
		},
	})
}

func UpdateCategory(c *fiber.Ctx) error{
	Id, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : "invalid id",
		})
	}
	category := models.Category{
		Id: uint(Id),
	}
	err =database.DB.First(&category).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "category not found",
		})
	}
	err = c.BodyParser(&category)
	if err != nil{
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when parsing category",
		})
	}
	err = database.DB.Model(&category).Updates(category).Error
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "wrong category info",
		})
	}
	return c.JSON(fiber.Map{
		"message" : "success",
		"payload" : fiber.Map{
			"category" : category,
		},
	})
}

func DeleteCategory(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid id",
		})
	}
	category := models.Category{
		Id: uint(id),
	}
	database.DB.Delete(&category)
	return nil
}