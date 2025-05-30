package todo

import (
	"bytes"
	"fmt"
	"text/tabwriter"
	"time"
)

// Domain godoc
//
// An interface that defines the behaviour for a todo item domain service struct.
type Domain interface {
	CreateItem(string) (Item, error)
	GetTabularItemList([]Item) (string, error)
	UpdateItemName(string, Item) (Item, error)
	CompleteItem(Item) (Item, error)
}

// defaultDomain godoc
//
// A structure which adheres to the Domain interface.
type defaultDomain struct{}

// NewDomain godoc
//
// Creates a new todo Domain instance.
func NewDomain() Domain {
	return &defaultDomain{}
}

// CreateItem godoc
//
// Creates a new todo Item instance and returns it.
//
// Returns nil and error if name is an empty string.
//
// Returns a new Item and nil on success.
func (d *defaultDomain) CreateItem(name string) (Item, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("CreateItem: `name` cannot be empty")
	}
	nowTime := time.Now()
	return NewItem(
		0,
		name,
		0,
		nowTime,
		nowTime,
	), nil
}

// GetTabularItemList godoc
//
// Returns a string representation of a tabular list of the items that are passed in.
//
// Returns empty string and error on error writing the list with the tab writer.
//
// Returns "No todo items..." and nil when items is an empty slice.
//
// Returns items in a tabular format and nil on success.
func (d *defaultDomain) GetTabularItemList(items []Item) (string, error) {
	if items == nil || len(items) == 0 {
		return fmt.Sprintf("No todo items...\n"), nil
	}

	var buffer bytes.Buffer

	padding := 4
	tabWidth := 4
	tw := tabwriter.NewWriter(&buffer, 0, tabWidth, padding, ' ', 0)

	// Write header to the tabWriter
	_, err := fmt.Fprintln(tw, "ID\tName\tLast Updated\tCreated\tIs Completed")
	if err != nil {
		return "", fmt.Errorf("GetTabularItemList: Error writing table header to tabWriter: %v", err)
	}

	// Write separator
	_, err = fmt.Fprintln(tw, "--\t----\t------------\t-------\t------------")
	if err != nil {
		return "", fmt.Errorf("GetTabularItemList: Error writing table header to tabWriter: %v", err)
	}

	for _, item := range items {
		completedIcon := "❌"
		if item.GetIsCompleted() == 1 {
			completedIcon = "✅"
		}
		formatting := "%d\t%s\t%s\t%s\t%s\n"
		_, err := fmt.Fprintf(
			tw,
			formatting,
			item.GetId(),
			item.GetName(),
			item.GetUpdatedAt().Format(time.DateTime),
			item.GetCreatedAt().Format(time.DateTime),
			completedIcon,
		)
		if err != nil {
			return "", fmt.Errorf(
				"GetTabularItemList: Error writing item %d: %v",
				item.GetId(), err,
			)
		}
	}

	if err := tw.Flush(); err != nil {
		return "", fmt.Errorf(
			"GetTabularItemList: Failed to flush tabWriter: %d", err,
		)
	}
	return buffer.String(), nil
}

// UpdateItemName godoc
//
// Updates the item name.
//
// Returns nil and error when name is empty or when the item is nil.
//
// Returns the updated item and nil on success.
func (d *defaultDomain) UpdateItemName(name string, item Item) (Item, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("UpdateItemName: name is empty")
	}
	if item == nil {
		return nil, fmt.Errorf("UpdateItemName: item is nil")
	}
	item.SetName(name)
	item.SetUpdatedAt(time.Now())
	return item, nil
}

// CompleteItem godoc
//
// Updates isCompleted on the item to 1 (true).
//
// Returns nil and error when the item is nil.
//
// Returns the updated item and nil on success.
func (d *defaultDomain) CompleteItem(item Item) (Item, error) {
	if item == nil {
		return nil, fmt.Errorf("CompleteItem: item is nil")
	}
	item.SetIsCompleted(1)
	item.SetUpdatedAt(time.Now())
	return item, nil
}
