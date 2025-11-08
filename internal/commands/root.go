package commands

import (
	"fmt"
	"os"

	"github.com/skaragianis/todo/internal/todo"
	"github.com/spf13/cobra"
)

var All bool

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Simple todo.txt cli application",
	Long:  "Todo.txt cli application designed around my wants",
	RunE: func(cmd *cobra.Command, args []string) error {
		ts := todo.NewService()

		todos, err := ts.ReadTodos("todo.txt")
		if err != nil {
			return err
		}

		fmt.Print(todos)

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&All, "all", "a", false, "display all todos, regardless if done")
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return nil
}
