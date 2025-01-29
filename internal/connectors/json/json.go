package json

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/mkhabelaj/todo/internal/todo"
)

type JsonFileConnector struct {
	FileName string
}

func (j *JsonFileConnector) Read() todo.TodoList {
	filePath := j.getFilePath()

	file, err := os.OpenFile(filePath, os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Failed to ", err)
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Failed to read file", err)
	}

	list := todo.TodoList{}

	if len(bytes) == 0 {
		return list
	}

	err = json.Unmarshal(bytes, &list)
	if err != nil {
		log.Fatal("Failed Unmarshal Json", err)
	}

	return list
}

func (j *JsonFileConnector) getFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Failed to get user home directory", err)
	}
	configDir := home + "/.todo"

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.Mkdir(configDir, 0755)
	}

	filePath := configDir + "/" + j.FileName
	return filePath
}

func (j *JsonFileConnector) Write(list *todo.TodoList) {
	filePath := j.getFilePath()
	data, err := json.Marshal(list)
	if err != nil {
		log.Fatal("Failed to MarshalIndent list:", err)
	}
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		log.Fatal("Failed to write file", err)
	}
}
