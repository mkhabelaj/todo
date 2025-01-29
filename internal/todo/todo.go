package todo

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
	"github.com/liamg/tml"

	"github.com/mkhabelaj/todo/internal/util"
)

type TodoItem struct {
	Info        string
	Completed   bool
	CreateAt    time.Time
	CompletedAT time.Time
}

type TodoList []TodoItem

type Todo struct {
	Connecter Connector
	list      *TodoList
}

func (t *Todo) Init() {
	if t.list == nil {
		t.list = &TodoList{}
	}
}

func (t *Todo) Load() {
	items := t.Connecter.Read()
	t.list = &items
}

func (t *Todo) Add(info string) {
	t.Init()
	newTodo := TodoItem{info, false, time.Now(), time.Time{}}
	*t.list = append(*t.list, newTodo)
	t.Save()
}

func (t *Todo) validateAndDecrementndex(index *int32) {
	todoList := *t.list

	if *index > int32(len(todoList)) || *index < 1 {
		log.Fatal("Invalid index")
	}

	*index--
}

func (t *Todo) Delete(index int32, save bool) {
	todoList := *t.list

	t.validateAndDecrementndex(&index)

	*t.list = append(todoList[:index], todoList[index+1:]...)

	if save {
		t.Save()
	}
}

func (t *Todo) DeleteMany(indexes []int32) {
	indexes = sortAndRemoveDuplicates(indexes)
	isFirst := true
	for _, index := range indexes {
		if isFirst {
			t.Delete(index, false)
			isFirst = false
			continue
		}

		// if a deletion already happened we need to decrement index to adjustj
		index--

		t.Delete(index, false)
	}

	t.Save()
}

func (t *Todo) CompleteMany(indexes []int32) {
	indexes = sortAndRemoveDuplicates(indexes)
	for _, index := range indexes {
		t.Complete(index, false)
	}

	t.Save()
}

func (t *Todo) Complete(index int32, save bool) {
	t.validateAndDecrementndex(&index)

	(*t.list)[index].Completed = true
	(*t.list)[index].CompletedAT = time.Now()

	t.Save()
}

func (t *Todo) Save() {
	t.Connecter.Write(t.list)
}

func (t *Todo) color(item string, color string) string {
	col := "<" + color + ">"
	return tml.Sprintf(col + item + col)
}

func (t *Todo) List() {
	list := *t.list

	if len(list) == 0 {
		log.Fatal("No todo's found")
	}
	for i, todo := range list {
		completed, completedAt := formatCompletionStatus(todo)

		if todo.Completed {
			completed = "true"

			fmt.Printf(
				"%v %v %v %v %v\n",
				i+1,
				completed,
				todo.Info,
				completedAt,
				todo.CreateAt.Format("2006-01-02 15:04:05"),
			)
			continue
		}
		completed = "false"
		fmt.Printf(
			"%v %v %v %v %v\n",
			i+1,
			completed,
			todo.Info,
			completedAt,
			todo.CreateAt.Format("2006-01-02 15:04:05"),
		)
	}
}

func (t *Todo) Table() {
	list := *t.list

	if len(list) == 0 {
		log.Fatal("No todo's found")
	}

	tabl := table.New(os.Stdout)
	tabl.SetHeaders("ID", "Completed", "Todo", "Completed At", "Created At")
	tabl.SetRowLines(false)

	for i, todo := range list {
		completed, completedAt := formatCompletionStatus(todo)

		if todo.Completed {
			completed = "✅"
			tabl.AddRow(t.color(strconv.Itoa(i+1), "green"),
				completed,
				t.color(todo.Info, "green"),
				t.color(completedAt, "green"),
				t.color(todo.CreateAt.Format("2006-01-02 15:04:05"), "green"),
			)
			continue
		}

		tabl.AddRow(
			strconv.Itoa(i+1),
			completed,
			todo.Info,
			completedAt,
			todo.CreateAt.Format("2006-01-02 15:04:05"),
		)
	}
	tabl.Render()
}

func formatCompletionStatus(todo TodoItem) (string, string) {
	completed := "-"
	completedAt := "-"

	if !todo.CompletedAT.IsZero() {
		completedAt = todo.CompletedAT.Format("2006-01-02 15:04:05")
	}
	return completed, completedAt
}

func sortAndRemoveDuplicates(ids []int32) []int32 {
	ids = util.RemoveDuplicatesInt(ids)
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
	return ids
}
