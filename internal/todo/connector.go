package todo

type Connector interface {
	Read() (TodoList, error)
	Write(list *TodoList) error
}
