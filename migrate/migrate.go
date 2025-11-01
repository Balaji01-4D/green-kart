package main

import (
	"fmt"
	"os"
	"shop-near-u/internal/models"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Migrate() {
	database := os.Getenv("DB_DATABASE")
	password := os.Getenv("DB_PASSWORD")
	username := os.Getenv("DB_USERNAME")
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	schema := os.Getenv("DB_SCHEMA")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s search_path=%s", host, username, password, database, port, sslmode, schema)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS postgis;")
	// Migrate the schema
	err = db.AutoMigrate(&models.User{})
	err = db.AutoMigrate(&models.Farmer{})
	err = db.AutoMigrate(&models.CatalogProduct{})
	err = db.AutoMigrate(&models.FarmerProduct{})
	err = db.AutoMigrate(&models.FarmerSubscription{})
	// Delivery-related models for hybrid logistics
	err = db.AutoMigrate(&models.DeliveryPartner{})
	err = db.AutoMigrate(&models.DeliveryOption{})
	err = db.AutoMigrate(&models.DeliveryRecord{})

	if err != nil {
		panic("failed to migrate database")
	}

	fmt.Println("Database migration completed successfully.")
}

func main() {
	Migrate()
}
