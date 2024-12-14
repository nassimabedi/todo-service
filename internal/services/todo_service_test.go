package services_test

import (
	"testing"
	"todo-service/internal/domain"
	"todo-service/internal/mocks"
	"todo-service/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestCreateTodoItem(t *testing.T) {
	mockRepo := mocks.NewMockTodoRepository()
	mockSQS := mocks.NewMockSQSClient()
	todoService := services.NewTodoService(mockRepo, mockSQS)

	todo := &domain.Todo{
		Description: "Test Todo",
		DueDate:     "2024-12-31T12:00:00Z",
		FileID:      "test-file-id",
	}

	err := todoService.CreateTodo(todo)
	assert.NoError(t, err)

	// Verify it was saved in the repository
	assert.Equal(t, 1, len(mockRepo.Todos))
	assert.Equal(t, "Test Todo", mockRepo.Todos[0].Description)

	// Verify it was sent to SQS
	assert.Equal(t, 1, len(mockSQS.Messages))
	assert.Equal(t, "Test Todo", mockSQS.Messages[0].Description)
}

