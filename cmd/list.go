package cmd

import (
	"fmt"
	"github.com/rykeroc/todo-cli/internal/data"
	"github.com/rykeroc/todo-cli/internal/modules/todo"
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todo items",
	Long:  "Displays a list of all existing todo items",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		dataSourceName := os.Getenv("DB_DATASOURCE_NAME")
		config := data.NewSqliteDatabaseConfig(dataSourceName)
		db := data.ConnectSqlDatabase(config)
		repository := todo.NewSqliteRepository(db)
		useCase := todo.NewUseCase(repository)

		err := useCase.List()
		if err != nil {
			log.Errorf("listCmd: %v", err)
			fmt.Println("An error occurred while listing todo items")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
