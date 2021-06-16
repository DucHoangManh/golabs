package controller

import (
	"fmt"

	"github.com/DucHoangManh/golabs/crud/model"
	repo "github.com/DucHoangManh/golabs/crud/repository"
	"github.com/gofiber/fiber/v2"
)
type BookController struct {}
func (bookController *BookController) Update(reviewI interface{}){
	review := reviewI.(*model.Review)
	sum := 0
	reviewCount := 0
	for _, rev := range repo.Reviews.GetAllReviews() {
		if review.BookId == rev.BookId{
			sum += rev.Rating
			reviewCount++
		}
	}
	var avgRating float32
	if len(repo.Reviews.GetAllReviews()) == 0{
		avgRating = 0
	}else{
		avgRating = float32(sum) / float32(reviewCount)
	}
	repo.Books.SetBookRating(review.BookId,avgRating)
}
func init() {
	bookController := &BookController{}
	repo.Reviews.Subscribe(bookController)
}
func GetAllBook(c *fiber.Ctx) error {
	return c.JSON(repo.Books.GetAllBooks())
}

func GetBookById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	book, err := repo.Books.FindBookById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(book)
}

func DeleteBookById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.Books.DeleteBookById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		repo.Reviews.DeleteReviewByBookId(int64(id))
		return c.SendString("delete successfully")
	}
}

func UpdateBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil{
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	book := new(model.Book)
	err = c.BodyParser(&book)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	repo.Books.UpdateBook(book,int64(id))
	return c.JSON(fiber.Map{
		"message": "update suceessfully",
		"book": book,
	})
}


func CreateBook(c *fiber.Ctx) error {
	book := new(model.Book)

	err := c.BodyParser(&book)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	bookId := repo.Books.CreateNewBook(book)
	return c.SendString(fmt.Sprintf("New book is created successfully with id = %d", bookId))

}
