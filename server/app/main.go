package main

import (
	"log"

  "go-buycar/internal/database"
  "go-buycar/api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.ConnectDB()

	router.AdvRoutes(app)
	router.CarRoutes(app)
  router.UserRoutes(app)
  router.SearchRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
