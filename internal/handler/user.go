package handler

import (
	"fmt"
	"strconv"

	"go-buycar/internal/database"
	"go-buycar/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == n
}

func validUser(id int, p string) bool {
	db := database.DB
	var user model.User
	db.First(&user, id)
	if user.Username == "" {
		return false
	}
	if !CheckPasswordHash(p, user.Password) {
		return false
	}
	return true
}

// GetUser get a user
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user model.User
	db.Find(&user, id)
	if user.Username == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

// CreateUser new user
func CreateUser(c *fiber.Ctx) error {
	type NewUser struct {
		ID			 int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	db := database.DB
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}

	user.Password = hash
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	newUser := NewUser{
		ID:				user.ID,
		Email:    user.Email,
		Username: user.Username,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}

// UpdateUser update user
func UpdateUser(c *fiber.Ctx) error {
	type UpdateUserInput struct {
		Name			  *string `json:"name"`
		Username    *string `gorm:"uniqueIndex;not null" json:"username"`
		Email       *string `gorm:"uniqueIndex;not null" json:"email"`
		Description *string `json:"description"`
		SecondName  *string `json:"second_name"`
		BirthDay		*string `json:"birthday"`
		ImageUrl    *string `json:"image_url"`
	}

	var uui UpdateUserInput
	if err := c.BodyParser(&uui); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
	}

	db := database.DB
		var user model.User
		if err := db.First(&user, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "User not found",
			})
		}

		if uui.Name != nil {
			user.Name = *uui.Name
		}
		if uui.Username != nil {
			user.Username = *uui.Username
		}
		if uui.Email != nil {
			user.Email = *uui.Email
		}
		if uui.Description != nil {
			user.Description = uui.Description
		}
		if uui.SecondName != nil {
			user.SecondName = uui.SecondName
		}
		if uui.ImageUrl != nil {
			user.ImageUrl = uui.ImageUrl
		}

		if err := db.Save(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to update user",
			})
		}

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "User updated successfully",
			"data":    user,
		})
}

// func RestorePassword(c* fiber.Ctx) error {}

func UpdateUserPassword(c* fiber.Ctx) error {
	type PasswordInput struct {
		CurrentPassword string  `json:"current_password"`
		NewPassword     string  `json:"new_password"`
		ID              *uint64 `json:"id"`
		Username        *string `json:"username"`
	}

	var input PasswordInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid input",
			"data":    err.Error(),
		})
	}

	if (input.ID == nil && input.Username == nil) || (input.ID != nil && input.Username != nil) {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Specify either ID or Username, but not both",
		})
	}

	var user model.User
	var err error

	db := database.DB
	if input.ID != nil {
		err = db.First(&user, *input.ID).Error
	} else {
		err = db.Where("username = ?", *input.Username).First(&user).Error
	}
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
		})
	}

	token := c.Locals("user").(*jwt.Token)
	if !validToken(token, fmt.Sprint(user.ID)) {
		return c.Status(403).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid token for user",
		})
	}

	if !validUser(user.ID, input.CurrentPassword) {
		return c.Status(403).JSON(fiber.Map{
			"status":  "error",
			"message": "Current password is incorrect",
		})
	}

	user.Password, err = hashPassword(input.NewPassword)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Password hashing failed",
		})
	}

	if err := db.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to update user",
			"data":    err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Password updated",
		"data":    user,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	type PasswordInput struct {
		Password			  string  `json:"password"`
		ID              *uint64 `json:"id"`
		Username        *string `json:"username"`
	}

	var input PasswordInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid input",
			"data":    err.Error(),
		})
	}

	if (input.ID == nil && input.Username == nil) || (input.ID != nil && input.Username != nil) {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Specify either ID or Username, but not both",
		})
	}

	var user model.User
	var err error

	db := database.DB
	if input.ID != nil {
		err = db.First(&user, *input.ID).Error
	} else {
		err = db.Where("username = ?", *input.Username).First(&user).Error
	}

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
		})
	}

	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, string(user.ID)) {
		return c.Status(403).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid token for user",
		})
	}

	if !validUser(user.ID, input.Password) {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "Not valid user",
			"data": nil})
	}

	db.Delete(&user)
	return c.JSON(fiber.Map{"status": "success", "message": "User successfully deleted", "data": nil})
}
