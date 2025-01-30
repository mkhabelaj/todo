/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mkhabelaj/todo/internal/connectors"
	"github.com/mkhabelaj/todo/internal/util"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Mark tasks as complete",
	Long: `Mark one or more tasks as complete by providing their IDs as arguments.

This command allows you to mark tasks as complete in your task list. 
For example:

check 1 2 3
[1|2|3] | check

This will mark tasks with IDs 1, 2, and 3 as complete. You can also pipe task IDs from another command.`,
	Run: func(cmd *cobra.Command, args []string) {
		TodoObj := connectors.GetConnectedTodo()
		err := TodoObj.Load()
		if err != nil {
			fmt.Println(err)
			return
		}
		pipeList := *util.ReadStdin()
		mergeList := append(pipeList, args...)

		intList, err := util.StrToint[int32](mergeList)
		if err != nil {
			fmt.Printf("Invalid input")
			return
		}
		err = TodoObj.CompleteMany(intList)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
