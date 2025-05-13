package router

import (
	"go-buycar/internal/handler"
	"go-buycar/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func AdvRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
  adverts := api.Group("/adverts", middleware.Protected())

  adverts.Get("/", handler.GetAllAdverts)
  adverts.Post("/post", handler.PostAdvert)

  adverts.Get("/:id<guid>", handler.GetAdvert)
  adverts.Get("/:car_name<alpha>", handler.GetAdvert)
  adverts.Get("/:owner_id<uint64>", handler.GetAdvertsByOwnerId)

  adverts.Patch("/:id<guid>", handler.UpdateAdvert)
  adverts.Delete("/:id<guid>", handler.DeleteAdvert)

	registerAdvCategoryRoutes(adverts)
}

func registerAdvCategoryRoutes(r fiber.Router) {
	categories := r.Group("/categories", middleware.Protected())

	categories.Post("/", handler.AddCategory)

  categories.Get("/", handler.GetCategories)
  categories.Get("/:name", handler.GetAdvertsByCategory)
}
