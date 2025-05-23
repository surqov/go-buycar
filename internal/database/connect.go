package database

import (
	"fmt"
	"strconv"

	"go-buycar/internal/config"
	"go-buycar/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic("failed to parse database port")
	}

  dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s \
    sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"),
    config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.User{}, &model.Car{}, &model.Advert{}, &model.Category{})
	fmt.Println("Database Migrated")
}
