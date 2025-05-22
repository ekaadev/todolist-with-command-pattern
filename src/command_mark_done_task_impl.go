package src

import "fmt"

// CONCRETE COMMAND: MARK DONE TASK

type CommandMarkAsDoneTaskImpl struct {
	todolist      *Todolist
	mark          int
	previousState string
}

func NewCommandMarkAsDoneTaskImpl(todolist *Todolist, mark int) Command {
	return &CommandMarkAsDoneTaskImpl{
		todolist: todolist,
		mark:     mark,
	}
}

func (c *CommandMarkAsDoneTaskImpl) Execute() {
	if c.mark < 0 || c.mark >= len(c.todolist.Todo) {
		fmt.Println("invalid mark")
		return
	}

	c.previousState = c.todolist.Todo[c.mark]

	c.todolist.MarkAsDone(c.mark)
}

func (c *CommandMarkAsDoneTaskImpl) Undo() {
	if c.mark < 0 || c.mark >= len(c.todolist.Todo) {
		fmt.Println("invalid mark")
		return
	}

	c.todolist.Todo[c.mark] = c.previousState
}

func (c *CommandMarkAsDoneTaskImpl) Redo() {
	c.todolist.MarkAsDone(c.mark)
}
