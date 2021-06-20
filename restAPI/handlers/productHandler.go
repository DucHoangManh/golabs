package handlers

import (
	"github.com/DucHoangManh/golabs/restAPI/database"
	"github.com/DucHoangManh/golabs/restAPI/models"
	"github.com/gofiber/fiber/v2"
)

func AllProducts(c *fiber.Ctx) error {
	var products []models.Product
	err := database.DB.Preload("Images").Find(&products).Error
	if err != nil {
		return err
	}
	return c.JSON(products)
}
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	err := c.BodyParser(&product)
	if err != nil {
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message": "can't parse request",
		})
	}
	err = database.DB.Create(&product).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" :err.Error(),
		})
	}
	err = createPrice(product.Id, product.Price)
	if err != nil{
		return err
	}
	database.DB.First(&product)
	return c.JSON(fiber.Map{
		"message" : "success",
		"product" : product,
	})
}

func GetProduct (c *fiber.Ctx) error {
	Id, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : "invalid id",
		})
	}
	product:= models.Product{
		Id: uint(Id),
	}
	database.DB.Preload("Images").First(&product)
	return c.JSON(fiber.Map{
		"message" : "success",
		"payload" : fiber.Map{
			"product" : product,
		},
	})
}

func UpdateProduct(c *fiber.Ctx) error{
	Id, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : "invalid id",
		})
	}
	product := models.Product{
		Id: uint(Id),
	}
	err =database.DB.First(&product).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "product not found",
		})
	}
	currentPrice := product.Price
	err = c.BodyParser(&product)
	if err != nil{
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when parsing product",
		})
	}
	err = database.DB.Model(&product).Updates(product).Error
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "wrong product info",
		})
	}
	if currentPrice != product.Price{
		err = createPrice(product.Id, product.Price)
		if err != nil{
			return err
		}
	}
	return c.JSON(fiber.Map{
		"message" : "success",
		"payload" : fiber.Map{
			"product" : product,
		},
	})
}

func DeleteProduct(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid id",
		})
	}
	product := models.Product{
		Id: uint(id),
	}
	database.DB.Delete(&product)
	return nil
}

func calculateRating(productId uint) error{
	var avgRating float32
	database.DB.Select("AVG(rating)").Where("product_id = ?", productId).Table("reviews").Find(&avgRating)
	product:= models.Product{
		Id: productId,
		Rating: avgRating,
	}
		err:= database.DB.Model(&product).Updates(product).Error
		if err != nil{
		return err
	}
	return nil
}