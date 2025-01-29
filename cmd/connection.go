package cmd

import (
	jconnector "github.com/mkhabelaj/todo/internal/connectors/json"
	"github.com/mkhabelaj/todo/internal/todo"
)

var (
	JsonConnector = jconnector.JsonFileConnector{FileName: "db.json"}
	TodoObj       = todo.Todo{Connecter: &JsonConnector}
)
