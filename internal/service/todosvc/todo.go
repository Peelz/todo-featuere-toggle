package todosvc

import (
	"context"
	"log"
	"monopeelz/pocff-go/internal/entities"
	"monopeelz/pocff-go/internal/ports"
)

type Service struct {
	ports.TodoRepository
	ports.FeatureFragmentProvider
}

func (s Service) SaveTodo(ctx context.Context) (entities.Todo, error) {
	result := entities.Todo{}
	// fResult := s.FeatureFragmentProvider.IsEnabledByStrategy(ctx, "save_todo")
	fResult := s.FeatureFragmentProvider.IsEnabledNumber(ctx, "save_todo")
	if fResult != 0 {
		result.Title = "save enabled"
		result.Note = fResult
		log.Println(fResult)
	} else {
		result.Title = "nothing"
	}
	return result, nil
}

func (s Service) ListTodo(ctx context.Context) (entities.Todo, error) {
	result := entities.Todo{}
	if s.FeatureFragmentProvider.IsEnabled(ctx, "list_todo") {
		result.Title = "list enabled"
	} else {
		result.Title = "off"
	}
	return result, nil
}

func NewService(
	todoRepository ports.TodoRepository,
	featureFragmentProvider ports.FeatureFragmentProvider,
) *Service {
	return &Service{TodoRepository: todoRepository, FeatureFragmentProvider: featureFragmentProvider}
}
