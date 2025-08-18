package config

import (
	"PracticalProject/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dns := "host=localhost user=postgres password=12345678 dbname=women_database port=5433 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Failed to connected to database!!!")
	}
	DB = database
	err = database.AutoMigrate(models.User{}, models.Product{}, models.Cart{}, models.Category{}, models.Inventory{}, models.Order{}, models.OrderItem{}, models.Payment{}, models.Promocode{})
	if err != nil {
		log.Fatal("Migrate failed ", err)
	}

	fmt.Println("Successfully connected to Database!!!")
}
