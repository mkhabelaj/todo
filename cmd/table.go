/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// tableCmd represents the table command
var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "Display tasks in a table format",
	Long: `Display tasks in a table format with options for sorting and filtering.
This command provides a structured view of tasks, allowing for easy management and review.`,
	Run: func(cmd *cobra.Command, args []string) {
		TodoObj.Load()
		TodoObj.Table()
	},
}

func init() {
	rootCmd.AddCommand(tableCmd)
}
