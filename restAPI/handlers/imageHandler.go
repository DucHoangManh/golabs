package handlers

import (
	"github.com/DucHoangManh/golabs/restAPI/database"
	"github.com/DucHoangManh/golabs/restAPI/models"
	"github.com/gofiber/fiber/v2"
)

func AllImages(c *fiber.Ctx) error {
	var images []models.Image
	err := database.DB.Find(&images).Error
	if err != nil {
		return err
	}
	return c.JSON(images)
}
func CreateImage(c *fiber.Ctx) error {
	var image models.Image
	err := c.BodyParser(&image)
	if err != nil {
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message": "can't parse request",
		})
	}
	err = database.DB.Create(&image).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" :err.Error(),
		})
	}
	database.DB.First(&image)
	return c.JSON(fiber.Map{
		"message" : "success",
		"image" : image,
	})
}

func GetImage (c *fiber.Ctx) error {
	Id, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : "invalid id",
		})
	}
	image:= models.Image{
		Id: uint(Id),
	}
	database.DB.Find(&image)
	return c.JSON(fiber.Map{
		"message" : "success",
		"payload" : fiber.Map{
			"image" : image,
		},
	})
}

func UpdateImage(c *fiber.Ctx) error{
	Id, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : "invalid id",
		})
	}
	image := models.Image{
		Id: uint(Id),
	}
	err =database.DB.First(&image).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "image not found",
		})
	}
	err = c.BodyParser(&image)
	if err != nil{
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when parsing image",
		})
	}
	err = database.DB.Model(&image).Updates(image).Error
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "wrong image info",
		})
	}
	return c.JSON(fiber.Map{
		"message" : "success",
		"payload" : fiber.Map{
			"image" : image,
		},
	})
}

func DeleteImage(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid id",
		})
	}
	image := models.Image{
		Id: uint(id),
	}
	database.DB.Delete(&image)
	return nil
}