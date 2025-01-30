package apple

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/mkhabelaj/todo/internal/todo"
	"github.com/mkhabelaj/todo/internal/util"
)

const (
	ID = "id"
)

type Reminders struct{}

func (t *Reminders) Add(todo todo.Todo, id int, save bool) error {
	if err := isApple(); err != nil {
		return err
	}
	task := todo.GetTask(int32(id))

	if todo.HasMeta(int32(id), ID) {
		return errors.New("Reminder already exists")
	}
	defautScript := fmt.Sprintf(
		`tell application "Reminders" to make new reminder with properties {name:"%s"}`,
		task.Info,
	)

	if !task.DueAt.IsZero() {
		dueDateString := task.DueAt.Format("1/2/2006 3:04:05 PM") // e.g. "2/20/2025 8:00:00 AM"

		defautScript = fmt.Sprintf(`
    tell application "Reminders"
        make new reminder with properties {name:"%s", due date:date "%s"}
    end tell
`, task.Info, dueDateString)
	}

	cmd := exec.Command("osascript", "-e", defautScript)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error running osascript: %s\n%s", err, string(output))
	}

	getId, err := getIdFromOutputBytes(output)
	if err != nil {
		return err
	}
	todo.UpdateMeta(int32(id), ID, getId, save)

	return nil
}

func (t *Reminders) AddMany(todo todo.Todo, infos []string, save bool) error {
	if err := isApple(); err != nil {
		return err
	}
	return errors.New("not implemented")
}

func (t *Reminders) Delete(todo todo.Todo, id int, save bool) error {
	if err := isApple(); err != nil {
		return err
	}
	return errors.New("not implemented")
}

func (t *Reminders) DeleteMany(todo todo.Todo, indexes []int32, save bool) error {
	if err := isApple(); err != nil {
		return err
	}
	return errors.New("not implemented")
}

func (t *Reminders) CompleteMany(todo todo.Todo, indexes []int32, save bool) error {
	if err := isApple(); err != nil {
		return err
	}
	return errors.New("not implemented")
}

func (t *Reminders) Complete(todo todo.Todo, id int, save bool) error {
	if err := isApple(); err != nil {
		return err
	}
	return errors.New("not implemented")
}

func (t *Reminders) List(todo todo.Todo) error {
	if err := isApple(); err != nil {
		return err
	}
	return errors.New("not implemented")
}

func isApple() error {
	if util.IsAppleComputer() {
		return nil
	}
	return errors.New("Current OS is not supported")
}

func getIdFromOutputBytes(bytes []byte) (string, error) {
	if len(bytes) == 0 {
		return "", errors.New("no bytes")
	}
	content := string(bytes)
	id := strings.Split(content, "//")[1]

	if id == "" {
		return "", errors.New("no id")
	}
	id = strings.TrimSpace(id)
	id = strings.Replace(id, "\n", "", -1)

	return id, nil
}
