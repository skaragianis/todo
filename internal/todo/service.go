package todo

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type TodoService struct{}

func NewService() *TodoService {
	return &TodoService{}
}

func (t *TodoService) ReadTodos(filename string) ([]Todo, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return t.loadTodos(file)
}

func (t *TodoService) SaveTodos(filename string, todos []Todo) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return t.writeTodos(file, todos)
}

func (t *TodoService) loadTodos(r io.Reader) ([]Todo, error) {
	scanner := bufio.NewScanner(r)

	todos := []Todo{}
	for scanner.Scan() {
		line := scanner.Text()

		todo, err := Parse(line)
		if err != nil {
			return nil, err
		}

		todos = append(todos, *todo)
	}

	return todos, nil
}

func (t *TodoService) writeTodos(w io.Writer, todos []Todo) error {
	for _, t := range todos {
		if _, err := fmt.Println(w, t); err != nil {
			return err
		}
	}

	return nil
}
