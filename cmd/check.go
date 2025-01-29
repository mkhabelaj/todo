/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"

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
		TodoObj.Load()
		pipeList := *util.ReadStdin()
		mergeList := append(pipeList, args...)
		intList := util.StrToint[int32](mergeList)
		TodoObj.CompleteMany(intList)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
