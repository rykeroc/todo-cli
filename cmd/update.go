package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:     `update <item id> "new name"`,
	Example: `todo update 1 "new name"`,
	Short:   "Update a todo item.",
	Long:    `Update the name of a todo item.`,
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		idToUpdate, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("Unable to update todo item.")
			fmt.Printf("'%s' is not a valid ID.\n", args[0])
			return
		}
		newName := args[1]

		updatedItemId, err := app.TodoUseCase.Update(idToUpdate, newName)
		if err != nil {
			log.Errorf("updateCmd: %v", err)
			fmt.Println("An error occurred while updating the todo item")
			return
		}
		if updatedItemId == -1 {
			fmt.Printf("No todo item exists with ID %d\n", idToUpdate)
			return
		}
		fmt.Println("Updated item")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
