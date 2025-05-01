package api

import (
	"github.com/gofiber/fiber/v2"
)

const var CurrentApiVersion = "1"

func CarRoutes(app *fiber.App) {
  api := app.Group("/api/v" + CurrentApiVersion)
  app.Get("/cars/:id", handlers.GetCarById)
}

func UserRoutes(app *fiber.App) {
  api := app.Group("/api/v" + CurrentApiVersion)
  app.Get("/users/:username<minLen(4)\>", handlers.GetUserInfo)
  app.Put("/users/:id<uint64\>", handlers.GetUserInfoer)
  app.Put("/users/:id<uint64\>", handlers.UpdateUser)

}
