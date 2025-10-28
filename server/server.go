package server

import (
	"gotodo/internal/todo"

	"github.com/labstack/echo/v4"
)

type Server struct {
	TodoService *todo.Service
}

func Run(todoService *todo.Service) {
	server := &Server{
		TodoService: todoService,
	}

	e := echo.New()

	e.HTTPErrorHandler = ErrorHandler

	server.registerRoutes(e)

	e.Logger.Fatal(e.Start(":5000"))
}

func (s *Server) registerRoutes(e *echo.Echo) {
	handler := &Handler{TodoService: s.TodoService}

	e.POST("/todos", handler.AddTodo)
	e.GET("/todos", handler.GetTodos)
	e.POST("/todos/:id/complete", handler.CompleteTodo)
	e.POST("/todos/:id/uncomplete", handler.UncompleteTodo)
	e.DELETE("/todos/:id", handler.DeleteTodo)
}
