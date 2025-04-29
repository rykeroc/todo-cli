package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// createCmd represents the `create` command
var createCmd = &cobra.Command{
	Use:     `create "<item name>"`,
	Example: `todo create "My new todo"`,
	Short:   "Create a todo item.",
	Long:    `Create a todo item with a specified name.`,
	Args:    cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		err := app.TodoUseCase.Create(args[0])
		if err != nil {
			log.Errorf("createCmd: %v\n", err)
			log.Fatalln("An error occurred while creating the todo item")
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
