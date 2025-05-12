package router

import (
	"go-buycar/internal/handler"
	"go-buycar/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func CarRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
  cars := api.Group("/cars")

  // Get All Cars
  cars.Get("/", middleware.Protected(), handler.GetAllCars)

  // Get CarInfo
  cars.Get("/:id<guid>", middleware.Protected(), handler.GetCar)
  cars.Get("/:car_name<alpha>", middleware.Protected(), handler.GetCar)
	cars.Get("/:owner_id<uint64>", middleware.Protected(), handler.GetCarsByOwnerId)

	// Manage Cars
	cars.Post("/add_new", handler.AddNewCar)
  cars.Patch("/:id<guid>", middleware.Protected(), handler.UpdateCar)
  cars.Delete("/:id<guid>", middleware.Protected(), handler.DeleteCar)

	registerCategoryRoutes(cars)
}

func registerCategoryRoutes(r fiber.Router) {
	categories := r.Group("/categories")

  categories.Get("/:name", middleware.Protected(), handler.GetCarsByCategory)
  categories.Get("/", middleware.Protected(), handler.GetCategoriesInfo)
  categories.Post("/", middleware.Protected(), handler.AddCategory)
}

