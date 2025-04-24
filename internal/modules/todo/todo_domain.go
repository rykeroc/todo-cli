package todo

import (
	"bytes"
	"fmt"
	"text/tabwriter"
	"time"
)

type Domain interface {
	CreateItem(string) (Item, error)
	GetItemList([]Item) (string, error)
	UpdateItem(string, Item) (Item, error)
}

type defaultDomain struct {
}

func NewDomain() Domain {
	return &defaultDomain{}
}

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

func (d *defaultDomain) GetItemList(items []Item) (string, error) {
	if items == nil || len(items) == 0 {
		return fmt.Sprintf("No todo items...\n"), nil
	}

	var buffer bytes.Buffer

	padding := 4
	tabWidth := 4
	tw := tabwriter.NewWriter(&buffer, 0, tabWidth, padding, ' ', 0)

	// Write header to the tabWriter
	_, err := fmt.Fprintln(tw, "ID\tName\tLast Updated\tCreation Date")
	if err != nil {
		return "", fmt.Errorf("GetItemList: Error writing table header to tabWriter: %v", err)
	}

	// Write separator
	_, err = fmt.Fprintln(tw, "--\t----\t------------\t-------------")
	if err != nil {
		return "", fmt.Errorf("GetItemList: Error writing table header to tabWriter: %v", err)
	}

	for _, item := range items {
		formatting := "%d\t%s\t%s\t%s\n"
		_, err := fmt.Fprintf(
			tw,
			formatting,
			item.GetId(),
			item.GetName(),
			item.GetUpdatedAt().Format(time.DateOnly),
			item.GetCreatedAt().Format(time.DateOnly),
		)
		if err != nil {
			return "", fmt.Errorf(
				"GetItemList: Error writing item %d: %v",
				item.GetId(), err,
			)
		}
	}

	if err := tw.Flush(); err != nil {
		return "", fmt.Errorf(
			"GetItemList: Failed to flush tabWriter: %d", err,
		)
	}
	return buffer.String(), nil
}

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
