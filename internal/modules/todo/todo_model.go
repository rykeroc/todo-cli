package todo

import (
	"database/sql"
	"fmt"
	"time"
)

// Item godoc
// Defines an interface for an item with getters and setters for encapsulation purposes.
type Item interface {
	GetId() int64
	GetName() string
	SetName(name string)
	GetDueAt() time.Time
	SetDueAt(time time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(time time.Time)
	GetCreatedAt() time.Time
}

// item godoc
// Defines an item structure.
// Implements the Item interface.
type item struct {
	id        int64
	name      string
	dueAt     time.Time
	updatedAt time.Time
	createdAt time.Time
}

// NewItem godoc
// Create a new instance of item which adheres to the Item interface.
func NewItem(
	id int64,
	name string,
	dueAt time.Time,
	updatedAt time.Time,
	createdAt time.Time,
) Item {
	return &item{
		id:        id,
		name:      name,
		dueAt:     dueAt,
		updatedAt: updatedAt,
		createdAt: createdAt,
	}
}

// NewItemFromRow godoc
// Create a new instance of item by scanning a sql.Rows struct.
func NewItemFromRow(rows *sql.Rows) (Item, error) {
	var item item
	var dueAtTimestamp, updatedAtTimestamp, createdAtTimestamp int64

	err := rows.Scan(&item.id, &item.name, &dueAtTimestamp, &updatedAtTimestamp, &createdAtTimestamp)
	if err != nil {
		return &item, fmt.Errorf("NewItemFromRow: %v", err)
	}

	item.dueAt = time.Unix(dueAtTimestamp, 0)
	item.updatedAt = time.Unix(updatedAtTimestamp, 0)
	item.createdAt = time.Unix(createdAtTimestamp, 0)

	return &item, nil
}

// GetId godoc
// Returns the item's ID.
func (item *item) GetId() int64 {
	return item.id
}

// GetName godoc
// Returns the item's name.
func (item *item) GetName() string {
	return item.name
}

// SetName godoc
// Sets the name of the item.
func (item *item) SetName(name string) {
	item.name = name
}

// GetDueAt godoc
// Returns the time that the item is due.
func (item *item) GetDueAt() time.Time {
	return item.dueAt
}

// SetDueAt godoc
// Sets the time that the item is due.
func (item *item) SetDueAt(time time.Time) {
	item.dueAt = time
}

// GetUpdatedAt godoc
// Returns the time that the item was last updated.
func (item *item) GetUpdatedAt() time.Time {
	return item.updatedAt
}

// SetUpdatedAt godoc
// Sets the updated at time of the item.
func (item *item) SetUpdatedAt(time time.Time) {
	item.updatedAt = time
}

// GetCreatedAt godoc
// Returns the time that the item was created.
func (item *item) GetCreatedAt() time.Time {
	return item.createdAt
}
