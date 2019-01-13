package firebase

import (
	"context"
	"github.com/kinnch/line-todo/models"
	"os"
)

type toDoRepository struct{}

type stubDoRepository struct{}

type iTodoRepository interface {
	Add(todo models.Todo) error
}

const (
	collectionName = "todo"
)

func NewTodoRepository() iTodoRepository {
	if os.Getenv("COMPONENT_TEST") != "" {
		return &stubDoRepository{}
	}
	return &toDoRepository{}
}

func (tr *toDoRepository) Add(todo models.Todo) error {
	ctx := context.Background()
	client, err := connect(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		todo.UserID: todo,
	})

	return err
}

func (tr *stubDoRepository) Add(todo models.Todo) error {
	return nil
}
