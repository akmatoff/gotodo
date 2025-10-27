package main

import (
	"fmt"
	"gotodo/todo"
	"os"
)

func main() {
	store := &todo.FileStore{FileName: "todos.json"}
	service := &todo.Service{Store: store}

	if len(os.Args) < 2 {
		fmt.Println("Expected add, list, complete, uncomplete or delete commands.")
		os.Exit(1)
	}

	switch os.Args[1] {

	case CmdAdd:
		task := os.Args[2]

		err := service.AddTodo(task)

		ExitOnError(err)

		fmt.Println("✅ Added todo:", task)
		fmt.Println()

		todos, _ := service.GetTodos()

		PrintTodosTable(todos)

	case CmdList:
		todos, err := service.GetTodos()

		ExitOnError(err)
		PrintTodosTable(todos)

	case CmdComplete:
		id, _ := ParseID(os.Args[2])
		err := service.CompleteTodo(id)

		ExitOnError(err)

		fmt.Println("✅ Completed todo with ID:", id)
		fmt.Println()

		todos, _ := service.GetTodos()

		PrintTodosTable(todos)

	case CmdUncomplete:
		id, _ := ParseID(os.Args[2])
		err := service.UncompleteTodo(id)

		ExitOnError(err)

		fmt.Println("✅ Uncompleted todo with ID:", id)
		fmt.Println()

		todos, _ := service.GetTodos()

		PrintTodosTable(todos)

	case CmdDelete:
		id, _ := ParseID(os.Args[2])
		err := service.DeleteTodo(id)

		ExitOnError(err)

		fmt.Println("✅ Deleted todo with ID:", id)
		fmt.Println()

		todos, _ := service.GetTodos()

		PrintTodosTable(todos)

	default:
		fmt.Println("Expected add, list, complete, uncomplete or delete commands.")
		os.Exit(1)

	}
}
