package main

import (
	"fmt"
	"github.com/rykeroc/todo-cli/cmd"
	"github.com/rykeroc/todo-cli/internal/config"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	logLevel, err := strconv.ParseInt(os.Getenv("LOG_LEVEL"), 0, 64)
	if err != nil {
		log.Fatalf("Invalid value for env var LOG_LEVEL: %s", err)
	}
	log.SetLevel(log.Level(logLevel))
	log.Debug(
		fmt.Sprintf("Log level: %d", logLevel),
	)

	cmd.Execute()
}
