/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package reminder

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mkhabelaj/todo/internal/api/apple"
	"github.com/mkhabelaj/todo/internal/connectors"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

		if appleB {
			reminder := apple.Reminders{}
			err = reminder.Add(TodoObj, 3, true)
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
