package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          int `gorm:"primaryKey"`
	Username    string `gorm:"uniqueIndex;not null" json:"username"`
	Email       string `gorm:"uniqueIndex;not null" json:"email"`
	Password    string `gorm:"not null" json:"password"`
	Description *string
	Name        string
	SecondName  *string
	BirthDay		*time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	ImageUrl    *string
}
