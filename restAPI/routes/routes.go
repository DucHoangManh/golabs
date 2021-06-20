package routes

import (
	"github.com/DucHoangManh/golabs/restAPI/handlers"
	"github.com/gofiber/fiber/v2"
)
func Setup(app *fiber.App){
		userApi := app.Group("api/user")
		userApi.Get("/", handlers.AllUsers)
		userApi.Post("/", handlers.CreateUser)
		userApi.Get("/:id", handlers.GetUser)
		userApi.Patch("/:id", handlers.UpdateUser)
		userApi.Put("/:id", handlers.PutUser)
		userApi.Delete("/:id", handlers.DeleteUser)

		categoryApi := app.Group("api/category")
		categoryApi.Get("/", handlers.AllCategories)
		categoryApi.Post("/", handlers.CreateCategory)
		categoryApi.Get("/:id", handlers.GetCategory)
		categoryApi.Patch("/:id", handlers.UpdateCategory)
		categoryApi.Delete("/:id", handlers.DeleteCategory)

		productApi := app.Group("api/product")
		productApi.Get("/", handlers.AllProducts)
		productApi.Post("/", handlers.CreateProduct)
		productApi.Get("/:id/price", handlers.GetPriceHistory)
		productApi.Get("/:id", handlers.GetProduct)
		productApi.Patch("/:id", handlers.UpdateProduct)
		productApi.Delete("/:id", handlers.DeleteProduct)
		

		imageApi := app.Group("api/image")
		imageApi.Get("/", handlers.AllImages)
		imageApi.Post("/", handlers.CreateImage)
		imageApi.Get("/:id", handlers.GetImage)
		imageApi.Patch("/:id", handlers.UpdateImage)
		imageApi.Delete("/:id", handlers.DeleteImage)

		reviewApi := app.Group("api/review")
		reviewApi.Get("/", handlers.AllReviews)
		reviewApi.Post("/", handlers.CreateReview)
		reviewApi.Get("/:id", handlers.GetReview)
		reviewApi.Patch("/:id", handlers.UpdateReview)
		reviewApi.Delete("/:id", handlers.DeleteReview)

		app.Get("/api/cart", handlers.AllCartItems)
		app.Post("/api/cart", handlers.CreateCartItem)
}
