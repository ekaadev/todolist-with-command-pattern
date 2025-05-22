package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todolist-with-command-pattern/src"
)

func main() {

	todo := src.Todolist{
		Todo: make([]string, 0),
	}

	history := src.HistoryCommand{
		History: make([]string, 0),
	}

	undo := src.UndoCommand{
		Undo: make([]src.Command, 0),
	}

	redo := src.RedoCommand{
		Redo: make([]src.Command, 0),
	}

	reader := bufio.NewReader(os.Stdin)
	var option, input string
	var err error
	var indexOfTodo int

loop:
	for {
		fmt.Println("=========================================")
		fmt.Printf("[HISTORY]: %v\n", history.History)
		fmt.Println("Todolist")
		for i := 0; i < len(todo.Todo); i++ {
			fmt.Printf("%v. %v\n", i+1, todo.Todo[i])
		}
		fmt.Println("=========================================")
		fmt.Println("1. Add todo")
		fmt.Println("2. Remove todo")
		fmt.Println("3. Mark as done todo")
		fmt.Println("u. Undo")
		fmt.Println("r. Redo")
		fmt.Println("x. Exit")
		fmt.Println("=========================================")
		fmt.Print("Choose one of the following options: ")
		option, err = reader.ReadString('\n')
		option = strings.TrimSpace(option)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch option {
		case "1":
			fmt.Print("Enter text: ")
			input, err = reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				break
			}
			add := src.NewCommandAddTaskImpl(&todo, strings.TrimSpace(input))
			add.Execute()
			undo.Undo = append(undo.Undo, add)
			redo.Redo = nil
			history.Add("ADD")
		case "2":
			fmt.Print("Choose one number of todo: ")
			input, err = reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				break
			}
			indexOfTodo, err = strconv.Atoi(strings.TrimSpace(input))
			if err != nil {
				fmt.Println(err)
				break
			}
			remove := src.NewCommandRemoveTaskImpl(&todo, indexOfTodo-1)
			remove.Execute()
			undo.Undo = append(undo.Undo, remove)
			redo.Redo = nil
			history.Add("REMOVE")
		case "3":
			fmt.Print("Choose one number of todo: ")
			input, err = reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				break
			}
			indexOfTodo, err = strconv.Atoi(strings.TrimSpace(input))
			if err != nil {
				fmt.Println(err)
				break
			}
			markAsDone := src.NewCommandMarkAsDoneTaskImpl(&todo, indexOfTodo-1)
			markAsDone.Execute()
			undo.Undo = append(undo.Undo, markAsDone)
			redo.Redo = nil
			history.Add("MARK")
		case "u":
			if len(undo.Undo) == 0 {
				fmt.Println("Undo is empty")
				break
			}
			// get last command
			cmd := undo.Undo[len(undo.Undo)-1]
			// execute last command
			cmd.Undo()
			// subtract stack
			undo.Undo = undo.Undo[:len(undo.Undo)-1]
			redo.Redo = append(redo.Redo, cmd)
			history.Add("UNDO")
		case "r":
			if len(redo.Redo) == 0 {
				fmt.Println("Redo is empty")
				break
			}
			// get last command
			cmd := redo.Redo[len(redo.Redo)-1]
			// execute last command
			cmd.Redo()
			// subtract stack
			redo.Redo = redo.Redo[:len(redo.Redo)-1]
			undo.Undo = append(undo.Undo, cmd)
			history.Add("REDO")
		case "x":
			break loop
		default:
			fmt.Println("Invalid option")
		}
	}

}
