package handlers

import (
	"github.com/DucHoangManh/golabs/restAPI/database"
	"github.com/DucHoangManh/golabs/restAPI/models"
	"github.com/gofiber/fiber/v2"
)

func AllUsers(c *fiber.Ctx) error {
	var users []models.User
	err := database.DB.Find(&users).Error
	if err != nil {
		return err
	}
	return c.JSON(users)
}
func CreateUser(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message": "can't parse request",
		})
	}
	err = database.DB.Create(&user).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" :err.Error(),
		})
	}
	database.DB.First(&user)
	return c.JSON(fiber.Map{
		"message" : "success",
		"user" : user,
	})
}

func GetUser (c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : "invalid id",
		})
	}
	user:= models.User{
		Id: uint(userId),
	}
	database.DB.Find(&user)
	return c.JSON(fiber.Map{
		"message" : "success",
		"payload" : fiber.Map{
			"user" : user,
		},
	})
}

func UpdateUser(c *fiber.Ctx) error{
	userId, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : "invalid id",
		})
	}
	user := models.User{
		Id: uint(userId),
	}
	err =database.DB.First(&user).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}
	err = c.BodyParser(&user)
	if err != nil{
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when parsing user",
		})
	}
	err = database.DB.Model(&user).Updates(user).Error
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "wrong person info",
		})
	}
	return c.JSON(fiber.Map{
		"message" : "success",
		"payload" : fiber.Map{
			"user" : user,
		},
	})
}
func PutUser(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : "invalid id",
		})
	}
	user := models.User{
		Id: uint(userId),
	}
	err =database.DB.First(&user).Error
	if err != nil {
		return CreateUser(c)
	}
	return UpdateUser(c)
}

func DeleteUser(c *fiber.Ctx) error{
	userId, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid id",
		})
	}
	user := models.User{
		Id: uint(userId),
	}
	database.DB.Delete(&user)
	return nil
}