package router

import (
	"go-buycar/internal/handler"
	"go-buycar/internal/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func UsersManagementRoute(app *fiber.App) {
	api := app.Group("/api", logger.New())
  user := api.Group("/users")

  // Auth
  user.Post("/login", handler.Login)
 
  // Registration
  user.Post("/registration", handler.CreateUser)

  // Get UserInfo
  user.Get("/:id", handler.GetUser)
  user.Get("/:username", handler.GetUser)

  // Update UserInfo
  user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
  user.Patch("/:username", middleware.Protected(), handler.UpdateUser)

  // User Deletion
  user.Delete("/:id", middleware.Protected(), handler.DeleteUser)
  user.Delete("/:username", middleware.Protected(), handler.DeleteUser)

	// // Cookies
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	// Create cookie
	// 	cookie := new(fiber.Cookie)
	// 	cookie.Name = "john"
	// 	cookie.Value = "doe"
	// 	cookie.Expires = time.Now().Add(24 * time.Hour)
	//
	// 	// Set cookie
	// 	c.Cookie(cookie)
	//
	//  return c.SendString("Cookie set!")
	// })
}
