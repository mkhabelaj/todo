package connectors

import (
	jconnector "github.com/mkhabelaj/todo/internal/connectors/json"
	"github.com/mkhabelaj/todo/internal/todo"
)

func GetConnectedTodo() todo.Todo {
	// TODO: should be configurable
	jsonConnector := jconnector.JsonFileConnector{FileName: "db.json"}

	todoObj := todo.Todo{Connecter: &jsonConnector}
	return todoObj
}
