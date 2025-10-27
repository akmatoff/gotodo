package main

import (
	"fmt"
	"gotodo/todo"
	"os"
	"strconv"
)

func main() {
	store := &todo.FileStore{FileName: "todos.json"}
	service := &todo.Service{Store: store}

	if len(os.Args) < 2 {
		fmt.Println("Expected add, list, complete, uncomplete or delete commands.")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "add":
		err := service.AddTodo(os.Args[2])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	case "list":
		todos, err := service.GetTodos()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, todo := range todos {
			fmt.Println(todo)
		}

	case "complete":
		id, _ := strconv.Atoi(os.Args[2])
		err := service.CompleteTodo(id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	case "uncomplete":
		id, _ := strconv.Atoi(os.Args[2])
		err := service.UncompleteTodo(id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	case "delete":
		id, _ := strconv.Atoi(os.Args[2])
		err := service.DeleteTodo(id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	default:
		fmt.Println("Expected add, list, complete, uncomplete or delete commands.")
		os.Exit(1)

	}
}
