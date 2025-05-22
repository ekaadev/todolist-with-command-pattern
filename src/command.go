package src

// COMMAND INTERFACE

type Command interface {
	Execute()
	Undo()
	Redo()
}
