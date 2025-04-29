package cmd

import (
	"fmt"
	"github.com/rykeroc/todo-cli/internal"
	"github.com/rykeroc/todo-cli/internal/data"
	"github.com/rykeroc/todo-cli/internal/modules/todo"
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/spf13/cobra"
)

type appComponents struct {
	TodoUseCase todo.UseCase
}

var helper data.SqlDatabaseHelper = nil
var app *appComponents = nil

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple CLI app for todo items.",
	Long:  "`todo` is appComponents simple CLI app for todo items.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Name() == "help" || cmd.Name() == "version" {
			return nil
		}
		log.Debugln("Running PersistentPreRunE")

		// Connect to database
		databaseFilename := fmt.Sprintf("%s.db", internal.AppName)
		helper = data.NewSqliteDatabaseHelper(databaseFilename)
		err := helper.Connect()
		if err != nil {
			log.Errorf("rootCmd: PersistentPreRunE: %v", err)
			return fmt.Errorf("an unexpected error occurred")
		}

		// Ensure database schema is initialized
		err = helper.InitializeSchema()
		if err != nil {
			log.Errorf("rootCmd: PersistentPreRunE: %v", err)
			return fmt.Errorf("an unexpected error occurred")
		}

		// Create the application structure
		db := helper.GetDatabase()
		if db == nil {
			log.Errorf("rootCmd: PersistentPreRunE: `db` from `helper` is uninitialized")
			return fmt.Errorf("an unexpected error occurred")
		}
		todoUseCase := todo.NewUseCase(
			todo.NewDomain(),
			todo.NewSqliteRepository(db),
		)
		app = &appComponents{
			todoUseCase,
		}

		log.Debugln("Completed PersistentPreRunE")
		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Name() == "help" || cmd.Name() == "version" {
			return nil
		}
		log.Debugln("Running PersistentPostRunE")

		// Return error if helper is not initialized
		if helper == nil {
			log.Errorf("rootCmd: PersistentPostRunE: `helper` is not initialized")
			return fmt.Errorf("an unexpected error occurred")
		}

		// Close database connection in database helper
		if err := helper.Close(); err != nil {
			log.Errorf("rootCmd: PersistentPostRunE: `helper` is not initialized")
			return fmt.Errorf("an unexpected error occurred")
		}

		app = nil

		log.Debugln("Completed PersistentPostRunE")
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
