package ports

import "monopeelz/pocff-go/internal/entities"

type TodoRepository interface {
	GetRepository() ([]entities.Todo, error)
}
