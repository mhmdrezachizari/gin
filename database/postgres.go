package database

import (
	"log"

	"backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=1234 dbname=shop port=5432 sslmode=disable TimeZone=Asia/Baku"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	DB = db

	err = DB.AutoMigrate(
		&models.Product{},
		&models.ProductImage{},
	)

	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}
}