package main

import (
	"fmt"
	"gotodo/internal/todo"
	"gotodo/pkg/constants"
	"gotodo/pkg/utils"
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

	case constants.CmdAdd:
		task := os.Args[2]

		_, err := service.AddTodo(task)

		utils.ExitOnError(err)

		fmt.Println("✅ Added todo:", task)
		fmt.Println()

		todos, _ := service.GetTodos()

		utils.PrintTodosTable(todos)

	case constants.CmdList:
		todos, err := service.GetTodos()

		utils.ExitOnError(err)
		utils.PrintTodosTable(todos)

	case constants.CmdComplete:
		id, _ := utils.ParseID(os.Args[2])
		err := service.CompleteTodo(id)

		utils.ExitOnError(err)

		fmt.Println("✅ Completed todo with ID:", id)
		fmt.Println()

		todos, _ := service.GetTodos()

		utils.PrintTodosTable(todos)

	case constants.CmdUncomplete:
		id, _ := utils.ParseID(os.Args[2])
		err := service.UncompleteTodo(id)

		utils.ExitOnError(err)

		fmt.Println("✅ Uncompleted todo with ID:", id)
		fmt.Println()

		todos, _ := service.GetTodos()

		utils.PrintTodosTable(todos)

	case constants.CmdDelete:
		id, _ := utils.ParseID(os.Args[2])
		err := service.DeleteTodo(id)

		utils.ExitOnError(err)

		fmt.Println("✅ Deleted todo with ID:", id)
		fmt.Println()

		todos, _ := service.GetTodos()

		utils.PrintTodosTable(todos)

	default:
		fmt.Println("Expected add, list, complete, uncomplete or delete commands.")
		os.Exit(1)

	}
}
