package apple

import (
	"errors"
	"fmt"
	"os/exec"

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
	task := (*todo.GetList())[id]

	if task.Meta[ID] != "" {
		return errors.New("Reminder already exists")
	}

	script := fmt.Sprintf(
		`tell application "Reminders" to make new reminder with properties {name:"%s", container:"%s"}`,
		task.Info,
		"Default",
	)
	cmd := exec.Command("osascript", "-e", script)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error running osascript: %s\n%s", err, string(output))
	}

	fmt.Println(string(output))

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
	if err := isApple(); err != nil {
		return err
	}
	if util.IsAppleComputer() {
		return nil
	}
	return errors.New("Current OS is not supported")
}
