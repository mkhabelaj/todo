package todo

type Connector interface {
	Read() TodoList
	Write(list *TodoList)
}
