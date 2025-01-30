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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long: `Delete a task by specifying its ID. You can provide multiple IDs to delete multiple items at once. 
  The command ensures that duplicate IDs are removed before processing. 
  Example usage:

  delete 1 2 3

  This will delete the todo items with IDs 1, 2, and 3.`,
	Run: func(cmd *cobra.Command, args []string) {
		TodoObj := connectors.GetConnectedTodo()
		err := TodoObj.Load()
		if err != nil {
			fmt.Println(err)
			return
		}
		pipeList := *util.ReadStdin()
		merge := append(args, pipeList...)

		ids := make([]int32, len(merge))
		ids, err = util.StrToint[int32](merge)
		if err != nil {
			fmt.Println("Invalid IDs, ensure they are numbers")
			return
		}

		err = TodoObj.DeleteMany(ids)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
