package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"todo-cli/internal/config"

	_ "github.com/mattn/go-sqlite3"
)

func getDatabase(dataSourceName string) *sql.DB {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connected to sqlite database: %s\n", dataSourceName)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Pinged database successfuly")

	return db
}

func closeDatabase(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}
	dataSourceName := os.Getenv("DB_DATASOURCE_NAME")

	db := getDatabase(dataSourceName)
	defer closeDatabase(db)

	fmt.Println("Hello world")
	return
}
