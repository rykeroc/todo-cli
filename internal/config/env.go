package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	// Check if the required environment variables are set
	var requiredVars = []string{
		"DB_DATASOURCE_NAME",
	}

	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			return fmt.Errorf("missing required environment variable: %s", v)
		}
	}

	return nil
}
