package router

import (
	"go-buyadverts/internal/handler"
	"go-buyadverts/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func AdvRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
  adverts := api.Group("/adverts")

  // Get All Adverts
  adverts.Get("/", middleware.Protected(), handler.GetAllAdverts)

  // Add NewAdvert
  adverts.Post("/add_new", handler.AddNewAdvert)

  // Get AdvertInfo
  adverts.Get("/:id<guid\>", middleware.Protected(), handler.GetAdvert)
  adverts.Get("/:car_name<alpha\>", middleware.Protected(), handler.GetAdvert)

  // Get Catogory list
  categories.Get("/", middleware.Protected(), handler.GetCategories)
  
  // Add New Category
  categories.Post("/", middleware.Protected(), handler.AddCategory)

  // Get Adverts from Category
  categories.Get("/:name", middleware.Protected(), handler.GetAdvertsByCategory)

  // Get Adverts by OwnerID
  adverts.Get("/:owner_id<uint64\>", middleware.Protected(), handler.GetAdvertsByOwnerId)

  // Update AdvertInfo
  adverts.Patch("/:id<guid\>", middleware.Protected(), handler.UpdateAdvert)

  // Advert Deletion
  adverts.Delete("/:id<guid\>", middleware.Protected(), handler.DeleteAdvert)

  //Get AllUserAdverts
  user.Get("/:id<uint64\>/adverts", middleware.Protected(), handler.GetUserAdverts)
}
