package data

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rykeroc/todo-cli/internal"
	log "github.com/sirupsen/logrus"
	"os"
)

//go:embed migrations/*.sql
var migrationsFs embed.FS

type SqlDatabaseHelper interface {
	Connect() error
	Close() error
	InitializeSchema() error
	GetDatabase() *sql.DB
}

type SqliteDatabaseHelper struct {
	DriverName       string
	DatabaseFilename string
	db               *sql.DB
}

func NewSqliteDatabaseHelper(databaseFilename string) SqlDatabaseHelper {
	driverName := "sqlite3"
	return &SqliteDatabaseHelper{
		DriverName:       driverName,
		DatabaseFilename: databaseFilename,
	}
}

func (s *SqliteDatabaseHelper) Connect() error {
	if len(s.DriverName) == 0 {
		return fmt.Errorf("Connect: invalid database config: Missing driver name")
	}

	if len(s.DatabaseFilename) == 0 {
		return fmt.Errorf("Connect: Invalid database config: Missing data source name")
	}

	if s.db == nil {
		databasePath, err := getDatabasePath(s.DatabaseFilename)
		if err != nil {
			return fmt.Errorf("Connect: %v", err)
		}

		if err := ensureDbIsCreated(databasePath); err != nil {
			return fmt.Errorf("Connect: %v", err)
		}
		dataSourceName := getDataSourceName(databasePath)

		s.db, err = sql.Open(s.DriverName, dataSourceName)
		if err != nil {
			return fmt.Errorf("Connect: %v", err)
		}
		log.Infof(
			"Connect: Opened %s database: %s\n",
			s.DriverName, s.DatabaseFilename,
		)
	}

	err := s.db.Ping()
	if err != nil {
		return fmt.Errorf("Connect: %v", err)
	}

	log.Infof("Connect: Pinged database successfully")
	return nil
}

func (s *SqliteDatabaseHelper) Close() error {
	if s.db == nil {
		return nil
	}
	err := s.db.Close()
	if err != nil {
		return fmt.Errorf("Close: %v", err)
	}
	log.Debugln("Close: Closed sqlite database connection")
	return nil
}

func (s *SqliteDatabaseHelper) InitializeSchema() error {
	if s.db == nil {
		return fmt.Errorf("InitializeSchema: db is not initialized")
	}
	migrationEntries, err := iofs.New(migrationsFs, "migrations")
	if err != nil {
		return fmt.Errorf("InitializeSchema: %v", err)
	}

	driver, err := sqlite.WithInstance(s.db, &sqlite.Config{})
	if err != nil {
		return fmt.Errorf("initializeDb: Failed to get sqlite driver instance: %v", err)
	}

	m, err := migrate.NewWithInstance(
		"iofs",
		migrationEntries,
		s.DatabaseFilename,
		driver,
	)
	if err != nil {
		return fmt.Errorf("InitializeSchema: Failed to get migrate instance: %v", err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("InitializeSchema: Failed run migrate up: %v", err)
	}

	return nil
}

func getDatabasePath(databaseName string) (string, error) {
	confDir, err := internal.GetAppConfigDir()
	if err != nil {
		return "", fmt.Errorf("getDatabasePath: %v", err)
	}
	return fmt.Sprintf("%s/%s", confDir, databaseName), nil
}

func getDataSourceName(path string) string {
	return fmt.Sprintf("file://%s", path)
}

func ensureDbIsCreated(dataSourceName string) error {
	_, err := os.Stat(dataSourceName)
	if nil == err {
		return nil
	}
	if os.IsNotExist(err) {
		_, err := os.Create(dataSourceName)
		if err != nil {
			return fmt.Errorf("ensureDbIsCreated: %v", err)
		}
		return nil
	}
	return fmt.Errorf("ensureDbIsCreated: %v", err)
}
func (s *SqliteDatabaseHelper) GetDatabase() *sql.DB {
	return s.db
}
