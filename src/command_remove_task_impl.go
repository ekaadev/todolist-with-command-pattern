package src

import "fmt"

// CONCRETE COMMAND: REMOVE TASK

type CommandRemoveTaskImpl struct {
	todolist      *Todolist
	mark          int
	previousState string
}

func NewCommandRemoveTaskImpl(todolist *Todolist, mark int) Command {
	return &CommandRemoveTaskImpl{
		todolist: todolist,
		mark:     mark,
	}
}

func (c *CommandRemoveTaskImpl) Execute() {
	if c.mark < 0 || c.mark >= len(c.todolist.Todo) {
		fmt.Println("invalid mark")
		return
	}

	c.previousState = c.todolist.Todo[c.mark]

	c.todolist.DeleteTodo(c.mark)
}

func (c *CommandRemoveTaskImpl) Undo() {

	c.todolist.InsertTodo(c.mark, c.previousState)
}

func (c *CommandRemoveTaskImpl) Redo() {
	c.todolist.DeleteTodo(c.mark)
}
