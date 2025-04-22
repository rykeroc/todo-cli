package testutils

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// TestFixture godoc
// A structure which contains dependencies for tests.
type TestFixture struct {
	Db *sql.DB
}

// SetupTestFixture godoc
// Creates an instance of TestFixture.
//
// Sets up an in memory SQLite database for integration tests.
func SetupTestFixture(t *testing.T) *TestFixture {
	db, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		t.Fatalf("SetupTestFixture: Failed to open in memory database connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		t.Fatalf("SetupTestFixture: Failed to ping database: %v", err)
	}

	if err = initializeSchema(db, t); err != nil {
		t.Fatalf("Failed to initialize schema: %v", err)
	}

	return &TestFixture{
		Db: db,
	}
}

// initializeSchema godoc
// Runs migrations on a SQLite database.
//
// Returns error on error, else nil
func initializeSchema(db *sql.DB, t *testing.T) error {
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		t.Fatalf("initializeSchema: Failed to get sqlite driver instance: %v", err)
	}

	migrationsDir := "file://./migrations"
	databaseName := "todo-cli"
	m, err := migrate.NewWithDatabaseInstance(
		migrationsDir,
		databaseName,
		driver,
	)
	if err != nil {
		t.Fatalf("initializeSchema: Failed to get migrate instance: %v", err)
	}

	err = m.Up()
	if err != nil {
		t.Fatalf("initializeSchema: Failed run migrate up: %v", err)
	}

	return nil
}

// CleanupTestFixture godoc
// Cleans up the dependencies that are created in SetupTestFixture.
func (tf *TestFixture) CleanupTestFixture() error {
	return tf.Db.Close()
}
