package ports

import "todo-service/internal/domain"

type SQSClient interface {
	SendMessage(queueURL string, message *domain.Todo) error
}

