package main

import (
	"fmt"
	"os"

	"github.com/skaragianis/todo/internal/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
