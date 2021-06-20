package handlers

import (
	"github.com/DucHoangManh/golabs/restAPI/database"
	"github.com/DucHoangManh/golabs/restAPI/models"
	"github.com/gofiber/fiber/v2"
)

func AllCartItems(c *fiber.Ctx) error {
	var cartItems []models.Cart
	err := database.DB.Preload("Product").Find(&cartItems).Error
	if err != nil {
		return err
	}
	return c.JSON(cartItems)
}
func CreateCartItem(c *fiber.Ctx) error {
	var cart models.Cart
	err := c.BodyParser(&cart)
	if err != nil {
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message": "can't parse request",
		})
	}
	err = database.DB.Create(&cart).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" :err.Error(),
		})
	}
	database.DB.First(&cart)
	return c.JSON(fiber.Map{
		"message" : "success",
		"cartItem" : cart,
	})
}



