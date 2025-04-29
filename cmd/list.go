package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todo items.",
	Long:  "Displays a list of all existing todo items.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		err := app.TodoUseCase.List()
		if err != nil {
			log.Errorf("listCmd: %v", err)
			fmt.Println("An error occurred while listing todo items")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
