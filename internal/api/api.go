package api

import "github.com/mkhabelaj/todo/internal/todo"

type Api interface {
	Add(todo todo.Todo, id int, save bool) error
	AddMany(todo todo.Todo, ids []int32, save bool) error
	Delete(todo todo.Todo, id int, save bool) error
	DeleteMany(todo todo.Todo, indexes []int32, save bool) error
	CompleteMany(todo todo.Todo, indexes []int32, save bool) error
	Complete(todo todo.Todo, id int, save bool) error
	List(todo todo.Todo) error
}
