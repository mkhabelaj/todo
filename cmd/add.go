package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mkhabelaj/todo/internal/connectors"
	"github.com/mkhabelaj/todo/internal/util"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new tasks to your todo list",
	Long: `Add new tasks to your todo list. You can provide tasks as arguments or through standard input.
For example:

To add tasks via arguments:
  todo add "Buy groceries" "Read a book"

To add tasks via standard input:
  echo "Go for a walk" | todo add`,
	Run: func(cmd *cobra.Command, args []string) {
		TodoObj := connectors.GetConnectedTodo()
		err := TodoObj.Load()
		if err != nil {
			fmt.Println(err)
			return
		}

		pipeList := *util.ReadStdin()
		mergeList := append(args, pipeList...)

		if err := TodoObj.AddMany(mergeList); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
