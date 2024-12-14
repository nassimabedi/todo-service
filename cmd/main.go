package main

import (
	"log"
	"todo-service/config"
	"todo-service/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Initialize database and services
	db := config.InitDB(cfg)
	s3Client := services.NewS3Service(cfg)
	sqsClient := services.NewSQSService(cfg)
	todoService := services.NewTodoService(db, s3Client, sqsClient)

	// Initialize router
	r := gin.Default()
	r.POST("/upload", todoService.UploadFile)
	r.POST("/todo", todoService.CreateTodo)

	log.Println("Starting server on :8080...")
	r.Run(":8080")
}

