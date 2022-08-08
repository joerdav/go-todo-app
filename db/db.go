package db

import "context"

type Todo struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	IsComplete bool   `json:"isComplete"`
}

type TodoStore interface {
	Update(context.Context, Todo) error
	LoadByUserID(context.Context, string) ([]Todo, error)
	Delete(context.Context, string) ([]Todo, error)
}
