package data

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

type SqlDatabaseConfig struct {
	driverName     string
	dataSourceName string
}

func NewSqlDatabaseConfig(driverName string, dataSourceName string) *SqlDatabaseConfig {
	return &SqlDatabaseConfig{
		driverName:     driverName,
		dataSourceName: dataSourceName,
	}
}

func NewSqliteDatabaseConfig(dataSourceName string) *SqlDatabaseConfig {
	driverName := "sqlite3"
	if len(dataSourceName) == 0 {
		log.Fatal("Unable to create SqlDatabaseConfig using sqlite3: Missing data source name")
	}
	return NewSqlDatabaseConfig(driverName, dataSourceName)
}

func ConnectSqlDatabase(config *SqlDatabaseConfig) *sql.DB {
	if len(config.driverName) == 0 {
		log.Fatal("Invalid database config: Missing driver name")
	}

	if len(config.dataSourceName) == 0 {
		log.Fatal("Invalid database config: Missing data source name")
	}

	db, err := sql.Open(config.driverName, config.dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(
		"Connected to %s database: %s\n",
		config.driverName, config.dataSourceName,
	)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Pinged database successfully")

	return db
}

func CloseSqlDatabase(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
