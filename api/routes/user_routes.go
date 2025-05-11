package router

import (
	"go-buycar/internal/handler"
	"go-buycar/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const (
	// username: 4 to 20 symbols
	// only chars, digs and '_'
	// case insensetive
	UserNameRegEx = `[a-zA-Z][a-zA-Z0-9_]{3,20}`
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	users := api.Group("/users")

	users.Post("/login", handler.Login)
	users.Post("/registration", handler.CreateUser)

	registerUserRoutesByID(users)
	registerUserRoutesByUsername(users)
}

func registerUserRoutesByID(r fiber.Router) {
	group := r.Group("/:id<uint64>", middleware.Protected())

	group.Get("/", handler.GetUser)
	group.Patch("/", handler.UpdateUser)
	group.Patch("/password", handler.UpdatePassword)
	group.Delete("/", handler.DeleteUser)
}

func registerUserRoutesByUsername(r fiber.Router) {
	group := r.Group("/:username<" + UserNameRegEx + ">", middleware.Protected())

	group.Get("/", handler.GetUser)
	group.Patch("/", handler.UpdateUser)
	group.Patch("/password", handler.UpdatePassword)
	group.Delete("/", handler.DeleteUser)
}
