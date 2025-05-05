package router

import (
	"go-buycar/internal/handler"
	"go-buycar/internal/middleware"

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
  user.Get("/:id<uint64\>", middleware.ProtectedUser(), handler.GetUser)
  user.Get("/:username<minLen(4)\>", middleware.ProtectedUser(), handler.GetUser)

  // Update UserInfo
  user.Patch("/:id<uint64\>", middleware.ProtectedUser(), handler.UpdateUser)
  user.Patch("/:username<minLen(4)\>", middleware.ProtectedUser(), handler.UpdateUser)

  // User Deletion
  user.Delete("/:id<uint64\>", middleware.ProtectedUser(), handler.DeleteUser)
  user.Delete("/:username<minLen(4)\>", mid", middleware.ProtectedUser(), handler.DeleteUser)
}

