package config

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	databaseURL := "postgres://postgres:2005b@localhost:5432/womens_shop?sslmode=disable"
	migrationPath := "file://db/migrations"

	m, err := migrate.New(migrationPath, databaseURL)
	if err != nil {
		log.Fatalf("Error creating migration insance : %v", err)
	}

	err = m.Up()
	if err != nil {
		if err.Error() == "no change" {
			fmt.Println("Database is up to date")
		} else {
			log.Fatalf("Error applying migrations: %v", err)
		}
	} else {
		fmt.Println("Migrations OK || Migrations applied successfully!!!")
	}
	dns := "host=localhost user=postgres password=2005b dbname=womens_shop port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Failed to connected to database!!!")
	}
	DB = database

	fmt.Println("GORM connected || Successfully connected to Database!!!")
}
