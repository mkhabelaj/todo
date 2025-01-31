package reminder

import (
	"github.com/spf13/cobra"
)

// reminderCmd represents the reminder command
var ReminderCmd = &cobra.Command{
	Use:   "reminder",
	Short: "Share local tasks with Apple Reminders and more",
	Long: `The 'reminder' command integrates local apps with external platforms, starting with Apple Reminders. 
Use subcommands for specific tasks. Future updates will support more platforms. Use 'help' for details.

Example usage:
  todo reminder add --apple 1 2 3 // add reminders for tasks 1, 2, and 3 to Apple Reminders
`,
}

func init() {
}
