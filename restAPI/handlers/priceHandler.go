package handlers

import (
	"github.com/DucHoangManh/golabs/restAPI/database"
	"github.com/DucHoangManh/golabs/restAPI/models"
	"github.com/gofiber/fiber/v2"
)

func createPrice(productId uint, value float64) error {
	price := models.Price{
		ProductId: productId,
		Value: value,
	}
	return database.DB.Create(&price).Error
}
func GetPriceHistory(c *fiber.Ctx) error{
	Id, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : "invalid id",
		})
	}
	var priceHistory []models.Price
	database.DB.Where("product_id = ?", Id).Table("prices").Find(&priceHistory)
	return c.JSON(priceHistory)
}