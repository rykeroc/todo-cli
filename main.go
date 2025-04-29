package main

import (
	"fmt"
	"github.com/rykeroc/todo-cli/cmd"
	"github.com/rykeroc/todo-cli/internal"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

// setLogLevel godoc
// Sets the log level for the application.
//
// Uses env variable `TODO_LOG_LEVEL` if present and valid.
//
// Defaults to Fatal log level.
func setLogLevel() {
	logLevel := log.FatalLevel
	envLogLevel := os.Getenv("TODO_LOG_LEVEL")
	if envLogLevel != "" {
		parsedLogLevel, err := strconv.ParseInt(
			envLogLevel,
			0,
			64,
		)
		if err != nil {
			log.Warnf("Invalid value for env var TODO_LOG_LEVEL: %s", err)
			log.SetLevel(internal.DefaultLogLevel)
			return
		}
		logLevel = log.Level(parsedLogLevel)
	}
	log.SetLevel(logLevel)
	log.Debug(
		fmt.Sprintf("Log level: %d", logLevel),
	)
}

// main godoc
// Entry point for the application.
func main() {
	setLogLevel()
	cmd.Execute()
}
