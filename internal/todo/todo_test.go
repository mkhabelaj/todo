package todo_test

import (
	"log"
	"os"
	"testing"

	connector "github.com/mkhabelaj/todo/internal/connectors/json"
	"github.com/mkhabelaj/todo/internal/todo"
)

func createTodoObj(filepath string) todo.Todo {
	conn := connector.JsonFileConnector{FilePath: filepath}
	return todo.Todo{Connecter: &conn}
}

func getFileBytes(filepath string) []byte {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func restoreFile(filepath string, fileBytes []byte) {
	err := os.WriteFile(filepath, fileBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func assertFileNotEmpty(filepath string, t *testing.T) {
	fileBytes := getFileBytes(filepath)
	if len(fileBytes) == 0 {
		t.Error("File is empty")
	}
}

func assertTodoListLength(todoList todo.TodoList, expectedLength int, t *testing.T) {
	if len(todoList) != expectedLength {
		t.Errorf("Expected list length to be %d, but got %d", expectedLength, len(todoList))
	}
}

func TestLoad(t *testing.T) {
	path := "testdata/test.json"
	// defer restoreFile(path, getFileBytes(path))
	assertFileNotEmpty(path, t)

	todoObj := createTodoObj(path)
	if todoObj.GetList() != nil {
		t.Error("Expected list to be nil before load")
	}
	err := todoObj.Load()
	if err != nil {
		t.Error(err)
	}

	if todoObj.GetList() == nil {
		t.Error("Expected list to not be nil after load")
	}
}

func TestAdd(t *testing.T) {
	path := "testdata/test.json"
	assertFileNotEmpty(path, t)
	defer restoreFile(path, getFileBytes(path))
	todoObj := createTodoObj(path)
	todoObj.Load()
	assertTodoListLength(*todoObj.GetList(), 6, t)
	err := todoObj.Add("Test", true)
	if err != nil {
		t.Error(err)
	}
	assertTodoListLength(*todoObj.GetList(), 7, t)
}

func TestDelete(t *testing.T) {
	path := "testdata/test.json"
	assertFileNotEmpty(path, t)
	defer restoreFile(path, getFileBytes(path))
	todoObj := createTodoObj(path)
	todoObj.Load()
	assertTodoListLength(*todoObj.GetList(), 6, t)
	err := todoObj.Delete(3, true)
	if err != nil {
		t.Error(err)
	}
	assertTodoListLength(*todoObj.GetList(), 5, t)
}

func TestComplete(t *testing.T) {
	path := "testdata/test.json"
	var zeroBaseIndexLocation int32 = 5
	var oneBaseIndexLocation int32 = zeroBaseIndexLocation + 1
	assertFileNotEmpty(path, t)
	defer restoreFile(path, getFileBytes(path))
	todoObj := createTodoObj(path)
	todoObj.Load()
	if (*todoObj.GetList())[zeroBaseIndexLocation].Completed != false {
		t.Errorf("Expected item at index %d to be not completed", zeroBaseIndexLocation)
	}
	err := todoObj.Complete(oneBaseIndexLocation, true)
	if err != nil {
		t.Error(err)
	}

	if (*todoObj.GetList())[zeroBaseIndexLocation].Completed != true {
		t.Errorf("Expected item at index %d to be not completed", zeroBaseIndexLocation)
	}
}

func TestAddMany(t *testing.T) {
	path := "testdata/test.json"
	assertFileNotEmpty(path, t)
	defer restoreFile(path, getFileBytes(path))
	todoObj := createTodoObj(path)
	todoObj.Load()
	assertTodoListLength(*todoObj.GetList(), 6, t)
	err := todoObj.AddMany([]string{"Test1", "Test2"})
	if err != nil {
		t.Error(err)
	}
	assertTodoListLength(*todoObj.GetList(), 8, t)
}

func TestCompleteMany(t *testing.T) {
	path := "testdata/test.json"
	assertFileNotEmpty(path, t)
	defer restoreFile(path, getFileBytes(path))
	todoObj := createTodoObj(path)
	todoObj.Load()

	if (*todoObj.GetList())[0].Completed != false {
		t.Errorf("Expected item at index %d to be not completed", 0)
	}
	if (*todoObj.GetList())[1].Completed != false {
		t.Errorf("Expected item at index %d to be not completed", 1)
	}
	err := todoObj.CompleteMany([]int32{1, 2})
	if err != nil {
		t.Error(err)
	}

	if (*todoObj.GetList())[0].Completed != true {
		t.Errorf("Expected item at index %d to be completed", 0)
	}
	if (*todoObj.GetList())[1].Completed != true {
		t.Errorf("Expected item at index %d to be completed", 1)
	}
}

func TestDeleteMany(t *testing.T) {
	path := "testdata/test.json"
	assertFileNotEmpty(path, t)
	defer restoreFile(path, getFileBytes(path))
	todoObj := createTodoObj(path)
	todoObj.Load()
	assertTodoListLength(*todoObj.GetList(), 6, t)
	err := todoObj.DeleteMany([]int32{1, 2})
	if err != nil {
		t.Error(err)
	}
	assertTodoListLength(*todoObj.GetList(), 4, t)
}
