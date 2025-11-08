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
		{"", Todo{projects: []string{}, contexts: []string{}, tags: map[string]string{}}, false},
		{"a simple todo", Todo{"a simple todo", false, 0, time.Time{}, time.Time{}, []string{}, []string{}, map[string]string{}}, false},
		{"A simple todo", Todo{"A simple todo", false, 0, time.Time{}, time.Time{}, []string{}, []string{}, map[string]string{}}, false},
		{"A simple todo done:true", Todo{"A simple todo done:true", false, 0, time.Time{}, time.Time{}, []string{}, []string{}, map[string]string{"done": "true"}}, false},
		{"A simple todo @some-context", Todo{"A simple todo @some-context", false, 0, time.Time{}, time.Time{}, []string{}, []string{"some-context"}, map[string]string{}}, false},
		{"A simple todo +some-project", Todo{"A simple todo +some-project", false, 0, time.Time{}, time.Time{}, []string{"some-project"}, []string{}, map[string]string{}}, false},
		{"A simple todo done:true @some-context +some-project", Todo{"A simple todo done:true @some-context +some-project", false, 0, time.Time{}, time.Time{}, []string{"some-project"}, []string{"some-context"}, map[string]string{"done": "true"}}, false},
		{"A simple todo done:true done2:false @some-context @some-context2 +some-project +some-project2", Todo{"A simple todo done:true done2:false @some-context @some-context2 +some-project +some-project2", false, 0, time.Time{}, time.Time{}, []string{"some-project", "some-project2"}, []string{"some-context", "some-context2"}, map[string]string{"done": "true", "done2": "false"}}, false},
		{"x A simple todo", Todo{"A simple todo", true, 0, time.Time{}, time.Time{}, []string{}, []string{}, map[string]string{}}, false},
		{"X A simple todo", Todo{"X A simple todo", false, 0, time.Time{}, time.Time{}, []string{}, []string{}, map[string]string{}}, false},
		{"(a) A simple todo", Todo{"(a) A simple todo", false, 0, time.Time{}, time.Time{}, []string{}, []string{}, map[string]string{}}, false},
		{"(A) A simple todo", Todo{"A simple todo", false, 'A', time.Time{}, time.Time{}, []string{}, []string{}, map[string]string{}}, false},
		{"(A A simple todo", Todo{"(A A simple todo", false, 0, time.Time{}, time.Time{}, []string{}, []string{}, map[string]string{}}, false},
		{"2025-02-01 A simple todo", Todo{"A simple todo", false, 0, time.Time{}, mustDate("2025-02-01"), []string{}, []string{}, map[string]string{}}, false},
		{"2025-42-01 A simple todo", Todo{}, true},
		{"x 2025-02-01 A simple todo", Todo{"A simple todo", true, 0, mustDate("2025-02-01"), time.Time{}, []string{}, []string{}, map[string]string{}}, false},
		{"x 2025-02-01 2025-06-10 A simple todo", Todo{"A simple todo", true, 0, mustDate("2025-02-01"), mustDate("2025-06-10"), []string{}, []string{}, map[string]string{}}, false},
		{"2025-02-01 A simple todo", Todo{"A simple todo", false, 0, time.Time{}, mustDate("2025-02-01"), []string{}, []string{}, map[string]string{}}, false},
		{"x", Todo{"x", false, 0, time.Time{}, time.Time{}, []string{}, []string{}, map[string]string{}}, false},
		{"x ", Todo{"x", false, 0, time.Time{}, time.Time{}, []string{}, []string{}, map[string]string{}}, true},
		{"2025-02-01", Todo{"2025-02-01", false, 0, time.Time{}, time.Time{}, []string{}, []string{}, map[string]string{}}, false},
		{"2025-02-01 ", Todo{}, true},
		{"2025-02-01 2025-03-02", Todo{"2025-03-02", false, 0, time.Time{}, mustDate("2025-02-01"), []string{}, []string{}, map[string]string{}}, false},
		{"2025-02-01 2025-03-02 ", Todo{}, true},
		{"2025-02-01 2025-40-02 ", Todo{}, true},
		{"2025-02-01 2025-03-02 A", Todo{}, true},
		{"x 2025-02-01 2025-06-10", Todo{"2025-06-10", true, 0, mustDate("2025-02-01"), time.Time{}, []string{}, []string{}, map[string]string{}}, false},
		{"x 2025-02-01 2025-06-10 ", Todo{}, true},

		// {"x 2016-05-20 2-16-04-30 a full example of a todo containing @context and a +tag plus some text on the end", Todo{}, false},
	}

	for _, test := range tests {
		result, err := Parse(test.in)
		if !test.expectErr && err == nil {
			if !reflect.DeepEqual(test.todo, *result) {
				t.Errorf("Parse(%s): Wanted: %v Got: %v", test.in, test.todo, *result)
			}
		} else {
			if err == nil && test.expectErr {
				t.Errorf("Parse(%s): Wanted: %v Got: %v", test.in, test.expectErr, err)
			} else if err != nil {
				t.Logf("Saw error: %v", err)
			}
		}
	}
}
