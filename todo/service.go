package todo

import (
	"fmt"
	"time"
)

type Service struct {
	Store Store
}

func (s *Service) AddTodo(text string) error {
	todos, err := s.Store.Load()

	if err != nil {
		return err
	}

	newTodo := Todo{
		Id:          len(todos) + 1,
		Text:        text,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	todos = append(todos, newTodo)

	return s.Store.Save(todos)
}

func (s *Service) GetTodos() ([]Todo, error) {
	return s.Store.Load()
}

func (s *Service) CompleteTodo(id int) error {
	todos, err := s.Store.Load()

	if err != nil {
		return err
	}

	found := false
	for i, todo := range todos {
		if todo.Id == id {
			todos[i].Complete()
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("Todo with ID %d not found", id)
	}

	return s.Store.Save(todos)
}

func (s *Service) UncompleteTodo(id int) error {
	todos, err := s.Store.Load()

	if err != nil {
		return err
	}

	for i, todo := range todos {
		if todo.Id == id {
			todos[i].Uncomplete()
		}
	}

	return s.Store.Save(todos)
}

func (s *Service) DeleteTodo(id int) error {
	todos, err := s.Store.Load()

	if err != nil {
		return err
	}

	found := false
	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("todo with id %d not found", id)
	}

	return s.Store.Save(todos)
}
