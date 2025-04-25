package cmd

import (
	"fmt"
	"github.com/rykeroc/todo-cli/internal/data"
	"github.com/rykeroc/todo-cli/internal/modules/todo"
	log "github.com/sirupsen/logrus"
	"os"
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

		dataSourceName := os.Getenv("DB_DATASOURCE_NAME")
		config := data.NewSqliteDatabaseConfig(dataSourceName)
		db := data.ConnectSqlDatabase(config)
		domain := todo.NewDomain()
		repository := todo.NewSqliteRepository(db)
		useCase := todo.NewUseCase(domain, repository)

		completedItemId, err := useCase.Complete(idToComplete)
		if err != nil {
			log.Errorf("completeCmd: %v", err)
			fmt.Println("An error occurred while completing the todo item")
			return
		}
		if completedItemId == -1 {
			fmt.Printf("No todo item exists with ID %d", idToComplete)
			return
		}
		fmt.Println("Completed item")
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
