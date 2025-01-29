/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"

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
		TodoObj.Load()

		if len(args) > 0 {
			for _, arg := range args {
				TodoObj.Add(arg)
			}
		}

		pipeList := *util.ReadStdin()

		if len(pipeList) > 0 {
			for _, arg := range pipeList {
				TodoObj.Add(arg)
			}
		}
		// stat, _ := os.Stdin.Stat()
		// if (stat.Mode() & os.ModeCharDevice) == 0 {
		// 	scanner := bufio.NewScanner(os.Stdin)
		//
		// 	for scanner.Scan() {
		// 		TodoObj.Add(scanner.Text())
		// 	}
		// 	if err := scanner.Err(); err != nil {
		// 		fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
		// 	}
		// }
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
