package todo

import (
	"errors"
	"time"
)

type Todo struct {
	description string
	completed   bool
	priority    byte
	completedAt time.Time
	createdAt   time.Time
}

func Parse(line string) (*Todo, error) {
	todo := &Todo{}

	if line == "" {
		return todo, nil
	}

	pos := 0

	if line[pos] == 'x' {
		todo.completed = true
		pos += 2
	} else if line[pos] == '(' {
		if pos+3 < len(line) {
			if line[pos+2] == ')' {
				candidate := line[pos+1]
				if candidate >= 'A' && candidate <= 'Z' {
					todo.priority = candidate
					pos += 4
				}
			}
		}
	}

	if pos >= len(line) {
		return nil, errors.New("you marked it completed but there's no detail for what?")
	}

	var date1, date2 time.Time
	var err error
	if line[pos] >= '0' && line[pos] <= '9' {
		date1, err = time.Parse("2006-01-02", line[pos:pos+10])
		if err != nil {
			return nil, errors.New("invalid date")
		}

		pos += 11
	}

	if pos >= len(line) {
		return nil, errors.New("you specified a date but then no detail for what?")
	}

	if line[pos] >= '0' && line[pos] <= '9' {
		// Assume a date
		date2, err = time.Parse("2006-01-02", line[pos:pos+10])
		if err != nil {
			return nil, errors.New("invalid date")
		}

		pos += 11
	}

	if pos >= len(line) {
		return nil, errors.New("you specified a creation date but no detail supporting why?")
	}

	if todo.completed {
		todo.completedAt = date1
		todo.createdAt = date2

	} else {
		todo.createdAt = date1
	}

	// The rest is the description
	todo.description = line[pos:]

	return todo, nil
}
