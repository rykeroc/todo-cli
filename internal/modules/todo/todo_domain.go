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
	UpdateItem(string, Item) (Item, error)
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
	_, err := fmt.Fprintln(tw, "ID\tName\tLast Updated\tCreated")
	if err != nil {
		return "", fmt.Errorf("GetTabularItemList: Error writing table header to tabWriter: %v", err)
	}

	// Write separator
	_, err = fmt.Fprintln(tw, "--\t----\t------------\t-------------")
	if err != nil {
		return "", fmt.Errorf("GetTabularItemList: Error writing table header to tabWriter: %v", err)
	}

	for _, item := range items {
		formatting := "%d\t%s\t%s\t%s\n"
		_, err := fmt.Fprintf(
			tw,
			formatting,
			item.GetId(),
			item.GetName(),
			item.GetUpdatedAt().Format(time.DateTime),
			item.GetCreatedAt().Format(time.DateTime),
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

// UpdateItem godoc
//
// Updates the item using the passed in arguments.
//
// Returns nil and error when name is empty or when the item is nil.
//
// Returns the updated item and nil on success.
func (d *defaultDomain) UpdateItem(name string, item Item) (Item, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("UpdateItem: name is empty")
	}
	if item == nil {
		return nil, fmt.Errorf("UpdateItem: item is nil")
	}
	item.SetName(name)
	item.SetUpdatedAt(time.Now())
	return item, nil
}
