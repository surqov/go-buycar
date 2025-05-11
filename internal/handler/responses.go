package handler

import (
	"github.com/gofiber/fiber/v2"
) 

type Response struct {
	status string
	message string
	data any
}

func ResponseToMap(lhs* Response) (fiber.Map, error) {
	return fiber.Map{
		"status": lhs.status, 
		"message": lhs.message, 
		"data": lhs.data, 
	}, nil
}
