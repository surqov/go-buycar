package router

import (
	"github.com/gofiber/fiber/v2"
)

var CurrentApiVersion string  = "1"

func CarRoutes(app *fiber.App) {
  // api := app.Group("/api/v" + CurrentApiVersion)
  app.Get("/cars/:id", func(c *fiber.Ctx) error {
    return c.SendString("id: " + c.Params("id"))
  })
}

// func UserRoutes(app *fiber.App) {
//   api := app.Group("/api/v" + CurrentApiVersion)
//   app.Get("/users/:username<minLen(4)\>", handlers.GetUserInfo)
//   app.Put("/users/:id<uint64\>", handlers.GetUserInfoer)
//   app.Put("/users/:id<uint64\>", handlers.UpdateUser)
// }
