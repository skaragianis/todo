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

const DateFormat = "2006-01-02"

func Parse(line string) (*Todo, error) {
	todo := &Todo{}

	if line == "" {
		return todo, nil
	}

	pos := 0

	switch line[pos] {
	case 'x':
		todo.completed = true
		line = line[2:]
	case '(':
		if pos+3 < len(line) {
			if line[pos+2] == ')' {
				candidate := line[pos+1]
				if candidate >= 'A' && candidate <= 'Z' {
					todo.priority = candidate
					line = line[4:]
				}
			}
		}
	}

	if len(line) == 0 {
		return nil, errors.New("you marked it complete, but what?")
	}

	if line[pos] >= '0' && line[pos] <= '9' {
		date, err := time.Parse(DateFormat, line[pos:pos+10])
		if err != nil {
			return nil, errors.New("invalid date")
		}

		if todo.completed {
			todo.completedAt = date
		} else {
			todo.createdAt = date
		}

		line = line[11:]
	}

	if len(line) == 0 {
		return nil, errors.New("you specified a date, but for what?")
	}

	if line[pos] >= '0' && line[pos] <= '9' {
		date, err := time.Parse(DateFormat, line[pos:pos+10])
		if err != nil {
			return nil, errors.New("invalid date")
		}

		if !todo.completed {
			return nil, errors.New("you provided a second date but the item isn't complete?")
		}

		todo.createdAt = date

		line = line[11:]
	}

	if len(line) == 0 {
		return nil, errors.New("you specified a creation date but for what?")
	}

	todo.description = line[pos:]

	return todo, nil
}
