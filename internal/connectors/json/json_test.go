package json

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/mkhabelaj/todo/internal/todo"
)

func TestReadFileEmpty(t *testing.T) {
	jsonReader := JsonFileConnector{FilePath: "testdata/emptyfile.json"}
	list, err := jsonReader.Read()
	if err != nil {
		t.Error("Expected no error, got", err)
	}

	if len(list) != 0 {
		t.Error("Expected empty list, got", list)
	}
}

func TestReadCreateIfNotExist(t *testing.T) {
	nonExistentFilePath := "testdata/new.json"

	// Ensure the file does not exist before the test
	if _, err := os.ReadFile(nonExistentFilePath); !os.IsNotExist(err) {
		t.Fatal("Expected no file to exist, but one was found:", err)
	}

	// Ensure the file is deleted after the test
	defer func() {
		if err := os.Remove(nonExistentFilePath); err != nil && !os.IsNotExist(err) {
			t.Log("Warning: Failed to remove test file:", err)
		}
	}()

	jsonReader := JsonFileConnector{FilePath: nonExistentFilePath}
	list, err := jsonReader.Read()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if len(list) != 0 {
		t.Fatalf("Expected an empty list but got: %v", list)
	}

	// Ensure the file now exists
	if _, err := os.ReadFile(nonExistentFilePath); os.IsNotExist(err) {
		t.Fatalf("Expected a file to exist, but none was found: %v", err)
	}
}

func TestReadFile(t *testing.T) {
	file := "testdata/test.json"
	jsonReader := JsonFileConnector{FilePath: file}
	list, err := jsonReader.Read()
	if err != nil {
		t.Error("Expected no error, got", err)
	}

	if len(list) == 5 {
		t.Error("Expected a list of 5, got", list)
	}
}

func TestWriteFile(t *testing.T) {
	file := "testdata/emptyfile.json"
	contents, err := os.ReadFile(file)

	if os.IsNotExist(err) {
		t.Error("Expected file to exist, but one was found:", err)
	}

	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if len(contents) > 0 {
		t.Error("Expected empty file, got", contents)
	}

	list := todo.TodoList{
		{
			Info:      "hello",
			Completed: false,
			CreateAt:  time.Now(),
		},
		{
			Info:      "world",
			Completed: false,
			CreateAt:  time.Now(),
		},
	}

	defer func() {
		// clear file
		err = os.WriteFile(file, []byte(""), 0644)
		if err != nil {
			fmt.Println("Failed to write file:", err)
		}
	}()

	jsonReader := JsonFileConnector{FilePath: file}
	err = jsonReader.Write(&list)
	if err != nil {
		t.Error("Expected no error, got", err)
	}

	data, err := jsonReader.Read()
	if err != nil {
		t.Error("Expected no error, got", err)
	}

	if len(data) != 2 {
		t.Error("Expected a list of 2, got", data)
	}
	if data[0].Info != "hello" {
		t.Error("Expected hello, got", data[0].Info)
	}
	if data[1].Info != "world" {
		t.Error("Expected world, got", data[1].Info)
	}
}
