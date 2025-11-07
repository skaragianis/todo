package todo

import (
	"reflect"
	"testing"
	"time"
)

func mustDate(s string) time.Time {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		panic(err)
	}

	return t
}

func TestParse(t *testing.T) {
	tests := []struct {
		in        string
		todo      Todo
		expectErr bool
	}{
		{"a simple todo", Todo{"a simple todo", false, 0, time.Time{}, time.Time{}}, false},
		{"A simple todo", Todo{"A simple todo", false, 0, time.Time{}, time.Time{}}, false},
		{"x A simple todo", Todo{"A simple todo", true, 0, time.Time{}, time.Time{}}, false},
		{"X A simple todo", Todo{"X A simple todo", false, 0, time.Time{}, time.Time{}}, false},
		{"(a) A simple todo", Todo{"(a) A simple todo", false, 0, time.Time{}, time.Time{}}, false},
		{"(A) A simple todo", Todo{"A simple todo", false, 'A', time.Time{}, time.Time{}}, false},
		{"(A A simple todo", Todo{"(A A simple todo", false, 0, time.Time{}, time.Time{}}, false},
		{"2025-02-01 A simple todo", Todo{"A simple todo", false, 0, time.Time{}, mustDate("2025-02-01")}, false},
		{"x 2025-02-01 A simple todo", Todo{"A simple todo", true, 0, mustDate("2025-02-01"), time.Time{}}, false},
		{"x 2025-02-01 2025-06-10 A simple todo", Todo{"A simple todo", true, 0, mustDate("2025-02-01"), mustDate("2025-06-10")}, false},
		{"2025-02-01 A simple todo", Todo{"A simple todo", false, 0, time.Time{}, mustDate("2025-02-01")}, false},

		// {"x 2016-05-20 2-16-04-30 a full example of a todo containing @context and a +tag plus some text on the end", Todo{}, false},
	}

	for _, test := range tests {
		result, err := Parse(test.in)
		if !test.expectErr && err == nil {
			if !reflect.DeepEqual(test.todo, *result) {
				t.Errorf("Parse(%s): Wanted: %v Got: %v", test.in, test.todo, *result)
			}
		} else {
			if err != nil {
				t.Errorf("Parse(%s): Wanted: %v Got: %v", test.in, test.expectErr, err)
			}
		}
	}
}
