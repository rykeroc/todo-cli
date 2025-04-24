package cmd

import (
	"fmt"
	"github.com/rykeroc/todo-cli/internal/data"
	"github.com/rykeroc/todo-cli/internal/modules/todo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

// createCmd represents the `create` command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a todo item.",
	Long: `Create a todo item by specifying a name:
	
	Example:
		todo create "My new todo"
	`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dataSourceName := os.Getenv("DB_DATASOURCE_NAME")
		config := data.NewSqliteDatabaseConfig(dataSourceName)
		db := data.ConnectSqlDatabase(config)
		repository := todo.NewSqliteRepository(db)
		useCase := todo.NewUseCase(repository)

		err := useCase.Create(args[0])
		if err != nil {
			log.Errorf("createCmd: %v", err)
			fmt.Println("An error occurred while creating the todo item")
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
