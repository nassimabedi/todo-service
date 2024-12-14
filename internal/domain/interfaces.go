package domain

import (
	"io"
)

// Todo represents a TodoItem entity
type Todo struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	FileID      string `json:"fileId"`
}

// S3Client defines methods for interacting with S3
type S3Client interface {
	UploadFile(bucketName, fileName string, file io.Reader) (string, error)
}

// SQSClient defines methods for interacting with SQS
type SQSClient interface {
	SendMessage(queueURL string, todo *Todo) error
}

// TodoService defines methods for Todo operations
type TodoService interface {
	UploadFile(file io.Reader, fileName string) (string, error)
	CreateTodo(description string, dueDate string, fileID string) (*Todo, error)
}

