package reminder

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mkhabelaj/todo/internal/api/apple"
	"github.com/mkhabelaj/todo/internal/connectors"
	"github.com/mkhabelaj/todo/internal/util"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new reminder",
	Long: `Add a new reminder from your list to a supported app. You can specify multiple IDs
to add multiple reminders at once. Use the --apple or -a flag to add
reminders specifically to the Apple Reminders app.

The workflow is to have Tasks in your todo list, then this command will convert them to Reminders
on the supported app. If the task has a due date, it will be added to the Reminders app as well.

Example usage:
  todo reminder add 1 2 3 --apple
  echo "1\n2\n3" | todo reminder add --apple

`,

	Run: func(cmd *cobra.Command, args []string) {
		appleB, err := cmd.Flags().GetBool("apple")
		if err != nil {
			fmt.Println(err)
		}

		TodoObj := connectors.GetConnectedTodo()
		err = TodoObj.Load()
		if err != nil {
			fmt.Println(err)
		}

		pipeList := *util.ReadStdin()
		merge := append(args, pipeList...)

		ids := make([]int32, len(merge))
		ids, err = util.StrToint[int32](merge)
		if err != nil {
			fmt.Println("Invalid IDs, ensure they are numbers")
			return
		}

		if appleB {
			reminder := apple.Reminders{}
			err = reminder.AddMany(TodoObj, ids, true)
			if err != nil {
				fmt.Println(err)
			}
			return
		}
		cmd.Help()
	},
}

func init() {
	ReminderCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("apple", "a", false, "Add an apple")
}
