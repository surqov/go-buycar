package handler

import (
	"strconv"

	"go-buycar/internal/database"
	"go-buycar/internal/model"

	"github.com/gofiber/fiber/v2"
)

func GetAllCars(c *fiber.Ctx) {
	db := database.DB
	var cars []model.Car
	if err := db.Find(&cars).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get cars",
		})
	}
	return c.JSON(cars)
}

// Add New Car to DB
func AddNewCar(c *fiber.Ctx) error {
	type NewCar struct {
		DoorsNum int `json:"doors_num"`
		SeatsNum int `json:"seats_num"`
	}

	db := database.DB
	car := new(model.Car)
	if err := c.BodyParser(car); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	if err := db.Create(&car).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create new car", "data": err})
	}

	newCar := NewCar{
		DoorNum:  user.DoorsNum,
		SeatsNum: user.SeatsNum,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created car", "data": newCar})
}
