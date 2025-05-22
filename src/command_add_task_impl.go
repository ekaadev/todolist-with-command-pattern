package src

// CONCRETE COMMAND: ADD TASK

type CommandAddTaskImpl struct {
	todolist *Todolist
	text     string
	index    int
}

func NewCommandAddTaskImpl(todolist *Todolist, text string) Command {
	return &CommandAddTaskImpl{
		todolist: todolist,
		text:     text,
	}
}

func (c *CommandAddTaskImpl) Execute() {
	c.todolist.AddTodo(c.text)
	c.index = len(c.todolist.Todo) - 1
}

func (c *CommandAddTaskImpl) Undo() {
	if c.index >= 0 && c.index < len(c.todolist.Todo) {
		c.todolist.DeleteTodo(c.index)
	}
}

func (c *CommandAddTaskImpl) Redo() {
	c.todolist.InsertTodo(c.index, c.text)
}
