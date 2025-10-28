package server

import (
	"fmt"
	"gotodo/internal/todo"
	"gotodo/pkg/utils"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	TodoService *todo.Service
}

func (h *Handler) GetTodos(c echo.Context) error {
	todos, err := h.TodoService.GetTodos()

	if err != nil {
		return err
	}

	return c.JSON(200, todos)
}

func (h *Handler) AddTodo(c echo.Context) error {
	var body todo.CreateTodoRequest

	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(500, "Internal server error")
	}

	fmt.Println(body)

	if len(body.Text) == 0 {
		return echo.NewHTTPError(400, "text is required")
	}

	newTodo, err := h.TodoService.AddTodo(body.Text)

	if err != nil {
		return echo.NewHTTPError(500, "Failed to create todo")
	}

	return c.JSON(201, map[string]interface{}{"message": "Todo created successfully", "data": newTodo})
}

func (h *Handler) CompleteTodo(c echo.Context) error {
	id, err := utils.ParseID(c.Param("id"))

	if err != nil {
		return err
	}

	if err := h.TodoService.CompleteTodo(id); err != nil {
		return echo.NewHTTPError(500, "Failed to complete todo")
	}

	return c.JSON(200, map[string]string{"message": "Todo completed successfully."})
}

func (h *Handler) UncompleteTodo(c echo.Context) error {
	id, err := utils.ParseID(c.Param("id"))

	if err != nil {
		return err
	}

	if err := h.TodoService.UncompleteTodo(id); err != nil {
		return err
	}

	return c.JSON(200, map[string]string{"message": "Uncompleted successfully!"})
}

func (h *Handler) DeleteTodo(c echo.Context) error {
	id, err := utils.ParseID(c.Param("id"))

	if err != nil {
		return err
	}

	if err := h.TodoService.DeleteTodo(id); err != nil {
		return err
	}

	return c.JSON(200, "Deleted successfully!")
}
