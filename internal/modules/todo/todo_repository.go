package todo

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
)

// Repository godoc
//
// Define a repository for a collection of Item.
type Repository interface {
	PersistItem(Item) (int64, error)
	FindAllItems() ([]Item, error)
	FindItemById(int64) (Item, error)
	UpdateItemById(Item) (int64, error)
	DeleteItemById(int64) (int64, error)
}

// sqliteRepository godoc
//
// Define a repository for a collection of Item that adheres to Repository.
type sqliteRepository struct {
	db *sql.DB
}

// NewSqliteRepository godoc
// Create a new instance of sqliteRepository that adheres to Repository.
func NewSqliteRepository(db *sql.DB) Repository {
	return &sqliteRepository{
		db: db,
	}
}

// tableName godoc
//
// Name for the database table which hold the items.
const tableName = "todos"

// PersistItem godoc
//
// Adds an Item to the database.
//
// Returns -1 and an error on error.
//
// Returns ID (Greater than 0) of inserted item and nil on success.
func (repo *sqliteRepository) PersistItem(itemToPersist Item) (int64, error) {
	if repo.db == nil {
		return -1, fmt.Errorf("PersistItem: database connection is nil")
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (displayName, updatedAt, createdAt) VALUES (?, ?, ?)",
		tableName,
	)
	result, err := repo.db.Exec(
		query,
		itemToPersist.GetName(),
		itemToPersist.GetUpdatedAt().Unix(),
		itemToPersist.GetCreatedAt().Unix(),
	)
	if err != nil {
		return -1, fmt.Errorf("PersistItem: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("PersistItem: %v", err)
	}

	return id, nil
}

// FindAllItems godoc
//
// Retrieves all items stored in the database table.
//
// Returns nil and error on error.
//
// Returns a slice containing Item instances and nil on success.
func (repo *sqliteRepository) FindAllItems() ([]Item, error) {
	if repo.db == nil {
		return nil, fmt.Errorf("FindAllItems: database connection is nil")
	}

	var result []Item
	query := fmt.Sprintf(
		"SELECT id, displayName, isCompleted, updatedAt, createdAt FROM %s ORDER BY isCompleted", tableName,
	)
	rows, err := repo.db.Query(query)
	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			if err == nil {
				// Return `closeErr` if `err` is not set already
				err = fmt.Errorf("FindAllItems failed to close rows: %w", closeErr)
			} else {
				// Log `closeErr` when `err` is already set
				log.Warnf("WARNING: FindAllItems failed to close rows (original error: %v): %v", err, closeErr)
			}
		}
	}(rows)
	if err != nil {
		return nil, fmt.Errorf("FindAllItems: %v", err)
	}

	result = []Item{}
	for rows.Next() {
		item, err := NewItemFromRow(rows)
		if err != nil {
			return nil, fmt.Errorf("FindAllItems: %v", err)
		}
		result = append(result, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("FindAllItems: %v", err)
	}

	return result, nil
}

// FindItemById godoc
//
// Get a persisted todo item by its ID.
//
// Returns nil and nil when no item is found.
//
// Returns nil and error on error.
//
// Returns the found Item and nil on success.
func (repo *sqliteRepository) FindItemById(id int64) (Item, error) {
	if id == 0 {
		return nil, nil
	}

	query := fmt.Sprintf(
		"SELECT id, displayName, isCompleted, updatedAt, createdAt FROM %s WHERE id = %d",
		tableName, id,
	)
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("FindItemById: %v", err)
	}
	// Close rows on exit
	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			if err == nil {
				// Return `closeErr` if `err` is not set already
				err = fmt.Errorf("FindItemById: Failed to close rows: %w", closeErr)
			} else {
				// Log `closeErr` when `err` is already set
				log.Warnf("WARNING: FindItemById: Failed to close rows (original error: %v): %v", err, closeErr)
			}
		}
	}(rows)

	if !rows.Next() {
		return nil, nil
	}

	item, err := NewItemFromRow(rows)
	if err != nil {
		return nil, fmt.Errorf("FindItemById: %v", err)
	}
	return item, nil
}

// UpdateItemById godoc
//
// Update an Item in the database table using its ID.
//
// Returns -1 and error on error.
//
// Returns number of updated rows and nil on success. If an item is updated the number of updated rows will be 1,
// else 0.
func (repo *sqliteRepository) UpdateItemById(itemToUpdate Item) (int64, error) {
	if repo.db == nil {
		return -1, fmt.Errorf("UpdateItemById: database connection is nil")
	}

	query := fmt.Sprintf(
		"UPDATE %s SET displayName = ?, updatedAt = ?, isCompleted = ? WHERE id = ?",
		tableName,
	)
	result, err := repo.db.Exec(
		query,
		itemToUpdate.GetName(),
		itemToUpdate.GetUpdatedAt().Unix(),
		itemToUpdate.GetIsCompleted(),
		itemToUpdate.GetId(),
	)
	if err != nil {
		return -1, fmt.Errorf("UpdateItemById: %v", err)
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("UpdateItemById: %v", err)
	}
	return rowCount, nil
}

// DeleteItemById godoc
//
// Delete an Item in the database table using its ID.
//
// Returns -1 and error on error.
//
// Returns number of deleted rows and nil on success. If an item is deleted the number of deleted rows will be 1,
// else 0.
func (repo *sqliteRepository) DeleteItemById(idToDelete int64) (int64, error) {
	if repo.db == nil {
		return -1, fmt.Errorf("DeleteItemById: database connection is nil")
	}

	query := fmt.Sprintf(
		"DELETE FROM %s WHERE id = ?",
		tableName,
	)
	result, err := repo.db.Exec(
		query,
		idToDelete,
	)
	if err != nil {
		return -1, fmt.Errorf("DeleteItemById: %v", err)
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("DeleteItemById: %v", err)
	}

	if rowCount == 1 {
		log.Infof("DeleteItemById: Successfullt deleted todo item with ID %d", idToDelete)
	} else {
		log.Infof("DeleteItemById: No item with ID %d", idToDelete)
	}
	return rowCount, nil
}
