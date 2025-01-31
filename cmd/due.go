package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/mkhabelaj/todo/internal/connectors"
	"github.com/mkhabelaj/todo/internal/util"
)

// dueCmd represents the due command
var dueCmd = &cobra.Command{
	Use:   "due",
	Short: "Set a due date for tasks",
	Long: `Set a due date for tasks using the specified date and time format.
For example:

todo due 1 2 3 --datetime "2025-12-31 23:59",
echo "1\n2\n3" | todo due --datetime "2025-12-31 23:59"`,
	Run: func(cmd *cobra.Command, args []string) {
		datetime, err := cmd.Flags().GetString("datetime")
		if err != nil {
			fmt.Println("Oops, something went wrong")
			return
		}
		if datetime == "" {
			fmt.Println("Please provide a valid date and time")
			cmd.Help()
			return
		}
		expectedFormat := "2006-01-02 15:04"
		parsedTime, err := time.Parse(expectedFormat, datetime)
		if err != nil {
			fmt.Println("Error: Invalid date format")
			return
		}

		TodoObj := connectors.GetConnectedTodo()
		err = TodoObj.Load()
		if err != nil {
			fmt.Println(err)
			return
		}
		pipeList := *util.ReadStdin()
		merge := append(args, pipeList...)
		convertedIds, err := util.StrToint[int32](merge)
		if err != nil {
			fmt.Println("Error: Please provide a valid ID")
			return
		}

		err = TodoObj.AddDueAtMany(convertedIds, parsedTime)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Task(s) due at", parsedTime, "has been added to your todo list")
	},
}

func init() {
	rootCmd.AddCommand(dueCmd)
	dueCmd.Flags().StringP("datetime", "t", "", "Due date in YYYY-MM-DD HH:MM format")
}
