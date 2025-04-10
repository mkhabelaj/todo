/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mkhabelaj/todo/internal/connectors"
)

// tableCmd represents the table command
var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "Display tasks in a table format",
	Long: `Display tasks in a table format with options for sorting and filtering.
This command provides a structured view of tasks, allowing for easy management and review.`,
	Run: func(cmd *cobra.Command, args []string) {
		TodoObj := connectors.GetConnectedTodo()
		err := TodoObj.Load()
		if err != nil {
			fmt.Println(err)
			return
		}
		TodoObj.Table()
	},
}

func init() {
	rootCmd.AddCommand(tableCmd)
}
