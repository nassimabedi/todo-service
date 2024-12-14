package ports

import "todo-service/internal/domain"

type Repository interface {
	SaveTodoItem(todo *domain.Todo) error
}

