package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/mkhabelaj/todo/cmd/reminder"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "todo",

	Short: "A simple CLI tool to manage your tasks efficiently",
	Long:  `Todo is a comprehensive command-line application designed to help you manage your tasks with ease and efficiency.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(reminder.ReminderCmd)
}
