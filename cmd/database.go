package cmd

import (
	"fmt"
	"log"
	"os"
	"persistence-service/model"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get raw DB:", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Ping failed:", err)
	}

	// Migrate tables
	if err := InitMigration(DB); err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Connected to PostgreSQL")
}

func InitMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Message{},
	)
}
