package utils

import (
	"fmt"
	"gotodo/internal/todo"
	"gotodo/pkg/constants"
	"os"
	"strconv"
	"strings"
)

func PrintTodosTable(todos []todo.Todo) {
	fmt.Printf("%-4s %-60s %-10s %-20s %-20s\n", "ID", "Text", "Completed", "CreatedAt", "CompletedAt")
	fmt.Println(strings.Repeat("-", 120))

	for _, t := range todos {
		status := "❌"
		if t.IsCompleted {
			status = "✅"
		}

		fmt.Printf("%-4d %-60s %-10s %-20s %-20s\n",
			t.Id,
			t.Text,
			status,
			t.CreatedAt.Format(constants.DATE_FORMAT),
			t.CompletedAt.Format(constants.DATE_FORMAT),
		)
	}
}

func ParseID(arg string) (int, error) {
	return strconv.Atoi(arg)
}

func ExitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
