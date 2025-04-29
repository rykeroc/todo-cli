package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     `remove <item id>`,
	Example: `todo remove 1`,
	Short:   "Delete a todo item.",
	Long:    "Delete an existing todo item by specifying it's ID.",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		idToDelete, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("Unable to remove todo item.")
			fmt.Printf("'%s' is not a valid ID.\n", args[0])
			return
		}

		deletedItemId, err := app.TodoUseCase.Remove(idToDelete)
		if err != nil {
			log.Errorf("removeCmd: %v", err)
			fmt.Println("An error occurred while removing the todo item")
			return
		}
		if deletedItemId == -1 {
			fmt.Printf("No todo item exists with ID %d", idToDelete)
			return
		}
		fmt.Println("Removed item")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
