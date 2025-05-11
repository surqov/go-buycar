package router

import (
	"go-buycar/internal/handler"
	"go-buycar/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// shortname: from 4 symbols
// only chars, digs and '_'
// case insensetive, means 'Gonb' == 'gonb'
const UserName_RE string = "^[a-z][a-z1-9_]{3,}$"

func CarRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
  car := api.Group("/cars")
  categories := car.Group("/categories")

  // Get All Cars
  car.Get("/", middleware.Protected(), handler.GetAllCars)

  // Add NewCar
  car.Post("/add_new", handler.AddNewCar)

  // Get CarInfo
  car.Get("/:id<guid\>", middleware.Protected(), handler.GetCar)
  car.Get("/:car_name<alpha\>", middleware.Protected(), handler.GetCar)

  // Get Cars by OwnerID
  car.Get("/:owner_id<uint64\>", middleware.Protected(), handler.GetCarsByOwnerId)

  // Update CarInfo
  car.Patch("/:id<guid\>", middleware.Protected(), handler.UpdateCar)

  // Car Deletion
  car.Delete("/:id<guid\>", middleware.Protected(), handler.DeleteCar)

  // Get Cars from Category
  categories.Get("/:name", middleware.Protected(), handler.GetCarsByCategory)
  //
  // Get Catogory list
  categories.Get("/", middleware.Protected(), handler.GetCategories)

  // Add New Category
  categories.Post("/", middleware.Protected(), handler.AddCategory)
}
