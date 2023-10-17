package appmodel

type AppState int

const (
	ListTodo AppState = iota
	AddTodo
)

type Model struct {
	State AppState
}

func NewModel() *Model {
	return &Model{
		State: ListTodo,
	}
}
