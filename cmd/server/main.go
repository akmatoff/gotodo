package main

import (
	"gotodo/internal/todo"
	"gotodo/server"
)

func main() {
	store := &todo.FileStore{FileName: "todos.json"}

	service := &todo.Service{Store: store}

	server.Run(service)
}
