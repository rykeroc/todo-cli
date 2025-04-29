package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:     "complete <item id>",
	Example: "todo complete 1",
	Short:   "Complete a todo item",
	Long:    "Complete a todo item by ID",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		idToComplete, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("Unable to complete todo item.")
			fmt.Printf("'%s' is not a valid ID.\n", args[0])
			return
		}

		completedItemId, err := app.TodoUseCase.Complete(idToComplete)
		if err != nil {
			log.Errorf("completeCmd: %v", err)
			log.Fatalln("An error occurred while completing the todo item")
		}
		if completedItemId == -1 {
			log.Fatalf("No todo item exists with ID %d\n", idToComplete)
		}
		fmt.Println("Completed item")
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
