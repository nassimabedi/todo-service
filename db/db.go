package db

import (
	"fmt"
	"log"
	"os"
	"todo-service/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes the database connection
func InitDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Run migrations
	err = db.AutoMigrate(&domain.Todo{})
	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	log.Println("Database connected and migrations applied.")
	return db
}

