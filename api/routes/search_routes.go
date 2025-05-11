package router

import (
	"go-buycar/internal/handler"
	"go-buycar/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)


func SearchRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
  car := api.Group("/adverts")

  // Get All Adverts
  car.Get("/", middleware.Protected(), handler.GetAllAdverts)

  // Add NewAdvert
  car.Post("/add_new", handler.AddNewAdvert)

  // Get AdvertInfo
  car.Get("/:id<guid\>", middleware.Protected(), handler.GetAdvert)
  car.Get("/:car_name<alpha\>", middleware.Protected(), handler.GetAdvert)

  // Get Catogory list
  categories.Get("/", middleware.Protected(), handler.GetCategories)
  
  // Add New Category
  categories.Post("/", middleware.Protected(), handler.AddCategory)

  // Get Adverts from Category
  categories.Get("/:name", middleware.Protected(), handler.GetAdvertsByCategory)

  // Get Adverts by OwnerID
  car.Get("/:owner_id<uint64\>", middleware.Protected(), handler.GetAdvertsByOwnerId)

  // Update AdvertInfo
  car.Patch("/:id<guid\>", middleware.Protected(), handler.UpdateAdvert)

  // Advert Deletion
  car.Delete("/:id<guid\>", middleware.Protected(), handler.DeleteAdvert)

  //Get AllUserAdverts
  user.Get("/:id<uint64\>/adverts", middleware.Protected(), handler.GetUserAdverts)
}
