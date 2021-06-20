package handlers

import (
	"github.com/DucHoangManh/golabs/restAPI/database"
	"github.com/DucHoangManh/golabs/restAPI/models"
	"github.com/gofiber/fiber/v2"
)

func AllReviews(c *fiber.Ctx) error {
	var reviews []models.Review
	err := database.DB.Find(&reviews).Error
	if err != nil {
		return err
	}
	return c.JSON(reviews)
}
func CreateReview(c *fiber.Ctx) error {
	var review models.Review
	err := c.BodyParser(&review)
	if err != nil {
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message": "can't parse request",
		})
	}
	err = database.DB.Create(&review).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" :err.Error(),
		})
	}
	err = calculateRating(review.ProductId)
	if err != nil{
		return err
	}
	database.DB.First(&review)
	return c.JSON(fiber.Map{
		"message" : "success",
		"review" : review,
	})
}

func GetReview (c *fiber.Ctx) error {
	Id, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : "invalid id",
		})
	}
	review:= models.Review{
		Id: uint(Id),
	}
	database.DB.First(&review)
	return c.JSON(fiber.Map{
		"message" : "success",
		"payload" : fiber.Map{
			"review" : review,
		},
	})
}

func UpdateReview(c *fiber.Ctx) error{
	Id, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : "invalid id",
		})
	}
	review := models.Review{
		Id: uint(Id),
	}
	err =database.DB.First(&review).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "review not found",
		})
	}
	err = c.BodyParser(&review)
	if err != nil{
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when parsing review",
		})
	}
	err = database.DB.Model(&review).Updates(review).Error
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "wrong review info",
		})
	}
	err = calculateRating(review.ProductId)
	if err != nil{
		return err
	}
	return c.JSON(fiber.Map{
		"message" : "success",
		"payload" : fiber.Map{
			"review" : review,
		},
	})
}

func DeleteReview(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid id",
		})
	}
	review := models.Review{
		Id: uint(id),
	}
	database.DB.Delete(&review)
	err = calculateRating(review.ProductId)
	if err != nil{
		return err
	}
	return nil
}