package services

import (
	"fmt"
	"log"
	"time"
	"todo-service/internal/domain"
	"todo-service/internal/ports"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoService struct {
	db        *gorm.DB
	s3Client  ports.S3Client
	sqsClient ports.SQSClient
}

func NewTodoService(db *gorm.DB, s3Client ports.S3Client, sqsClient ports.SQSClient) *TodoService {
	return &TodoService{
		db:        db,
		s3Client:  s3Client,
		sqsClient: sqsClient,
	}
}

func (t *TodoService) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "No file provided"})
		return
	}

	// Validate file type and size
	if file.Size > 10<<20 { // Limit: 10 MB
		c.JSON(400, gin.H{"error": "File too large"})
		return
	}

	openedFile, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "Unable to open file"})
		return
	}
	defer openedFile.Close()

	fileKey := fmt.Sprintf("%s-%s", uuid.New().String(), file.Filename)
	fileURL, err := t.s3Client.UploadFile("bucket-name", fileKey, openedFile)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to upload file to S3"})
		return
	}

	c.JSON(200, gin.H{"fileId": fileKey, "fileUrl": fileURL})
}

func (t *TodoService) CreateTodo(c *gin.Context) {
	var req struct {
		Description string `json:"description"`
		DueDate     string `json:"due_date"`
		FileID      string `json:"file_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Validate input
	if req.Description == "" || req.DueDate == "" {
		c.JSON(400, gin.H{"error": "Missing required fields"})
		return
	}

	dueDate, err := time.Parse(time.RFC3339, req.DueDate)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid due date format"})
		return
	}

	// Create TodoItem
	todo := &domain.Todo{
		ID:          uuid.New().String(),
		Description: req.Description,
		DueDate:     dueDate.Format(time.RFC3339),
		FileID:      req.FileID,
	}

	// Save to DB
	if err := t.db.Create(todo).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to save todo item"})
		return
	}

	// Send to SQS
	if err := t.sqsClient.SendMessage("queue-url", todo); err != nil {
		c.JSON(500, gin.H{"error": "Failed to send message to SQS"})
		return
	}

	c.JSON(200, todo)
}

