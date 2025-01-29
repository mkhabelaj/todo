package json

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"

	"github.com/mkhabelaj/todo/internal/todo"
)

type JsonFileConnector struct {
	FileName string
	FilePath string
}

func (j *JsonFileConnector) Read() (todo.TodoList, error) {
	filePath := j.getFilePath()

	file, err := os.OpenFile(filePath, os.O_CREATE, 0644)
	if err != nil {
		err = errors.New("Failed to open file: " + err.Error())
		return nil, err
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		err = errors.New("Failed to read file: " + err.Error())
		return nil, err
	}

	list := todo.TodoList{}

	if len(bytes) == 0 {
		return list, nil
	}

	err = json.Unmarshal(bytes, &list)
	if err != nil {
		err = errors.New("Failed Unmarshal Json: " + err.Error())
		return nil, err
	}

	return list, nil
}

func (j *JsonFileConnector) getFilePath() string {
	if j.FilePath != "" {
		return j.FilePath
	}
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

func (j *JsonFileConnector) Write(list *todo.TodoList) error {
	filePath := j.getFilePath()
	data, err := json.Marshal(list)
	if err != nil {
		return errors.New("Failed to MarshalIndent list: " + err.Error())
	}
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		log.Fatal("Failed to write file", err)
	}
	return nil
}
