package router

import (
	"go-buycar/internal/handler"
	"go-buycar/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const UserNameRegEx = "[a-zA-Z][a-zA-Z0-9_]{3,20}"

func UserRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	users := api.Group("/users")

	api.Post("/register", handler.CreateUser)
	api.Post("/login", handler.Login)
	api.Post("/login/restore", handler.RestorePassword)

	registerUserRoutesByID(users)
	registerUserRoutesByUsername(users)
}

func registerUserRoutesByID(r fiber.Router) {
	users_id := r.Group("/:id<uint64>", middleware.Protected())

	users_id.Get("/", handler.GetUser)
	users_id.Patch("/", handler.UpdateUser)
	users_id.Post("/password", handler.UpdateUserPassword)
	users_id.Delete("/", handler.DeleteUser)

	users_id.Get("/adverts", handler.GetUserAdverts)
}

func registerUserRoutesByUsername(r fiber.Router) {
	users_username := r.Group("/:username<" + UserNameRegEx + ">", middleware.Protected())

	users_username.Get("/", handler.GetUser)
	users_username.Patch("/", handler.UpdateUser)
	users_username.Post("/password", handler.UpdateUserPassword)
	users_username.Delete("/", handler.DeleteUser)

	users_username.Get("/adverts", handler.GetUserAdverts)
}
