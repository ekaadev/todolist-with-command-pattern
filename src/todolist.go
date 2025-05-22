package src

import "fmt"

// CONCRETE RECEIVER

type Todolist struct {
	Todo []string
}

func (t *Todolist) AddTodo(text string) {
	t.Todo = append(t.Todo, text)
}

func (t *Todolist) DeleteTodo(index int) {
	if len(t.Todo) == 0 {
		fmt.Println("todolist is empty")
		return
	} else if index < 0 || index >= len(t.Todo) {
		fmt.Println("invalid index")
		return
	}

	t.Todo = append(t.Todo[:index], t.Todo[index+1:]...)
}

func (t *Todolist) MarkAsDone(index int) {
	if len(t.Todo) == 0 {
		fmt.Println("todolist is empty")
		return
	} else if index < 0 || index >= len(t.Todo) {
		fmt.Println("invalid index")
		return
	}

	t.Todo[index] = "DONE"
}

func (t *Todolist) InsertTodo(index int, text string) {
	if index < 0 || index > len(t.Todo) {
		fmt.Println("invalid index")
		return
	}

	t.Todo = append(t.Todo[:index], append([]string{text}, t.Todo[index:]...)...)
}
