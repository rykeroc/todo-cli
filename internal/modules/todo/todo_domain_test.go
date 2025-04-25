package todo

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

var domain = NewDomain()

func TestDefaultDomain_CreateItem(t *testing.T) {
	t.Run("should create new item", func(t *testing.T) {
		name := "new item"

		item, err := domain.CreateItem(name)

		assert.NoError(t, err)
		assert.Equal(t, name, item.GetName())
	})

	t.Run("should return error because of empty name", func(t *testing.T) {
		name := ""

		item, err := domain.CreateItem(name)

		assert.Error(t, err)
		assert.Nil(t, item)
	})
}

func TestDefaultDomain_GetItemList(t *testing.T) {
	t.Run("should return tabular list", func(t *testing.T) {
		nowTime := time.Now()
		items := []Item{
			NewItem(1, "item 1", 0, nowTime, nowTime),
		}
		result, err := domain.GetTabularItemList(items)
		assert.NoError(t, err)
		assert.NotEmpty(t, result)

		assert.Contains(
			t, result, strconv.FormatInt(items[0].GetId(), 10),
		)
		assert.Contains(
			t, result, items[0].GetName(),
		)
		assert.Contains(
			t, result, items[0].GetUpdatedAt().Format(time.DateOnly),
		)
		assert.Contains(
			t, result, items[0].GetCreatedAt().Format(time.DateOnly),
		)
	})

	t.Run("should return 'No todo items...' when no items is nil", func(t *testing.T) {
		result, err := domain.GetTabularItemList(nil)

		assert.NoError(t, err)
		assert.Contains(t, result, "No todo items...")
	})

	t.Run("should return 'No todo items...' when no items is empty", func(t *testing.T) {
		result, err := domain.GetTabularItemList([]Item{})

		assert.NoError(t, err)
		assert.Contains(t, result, "No todo items...")
	})
}

func TestDefaultDomain_UpdateItem(t *testing.T) {
	t.Run("should update name and updated time in item", func(t *testing.T) {
		initName := "init name"
		initCreated := time.Now()
		initUpdated := time.Now()

		item := NewItem(
			0, initName, 0, initCreated, initUpdated,
		)

		newName := "new name"

		item, err := domain.UpdateItemName(newName, item)

		assert.NoError(t, err)
		assert.Equal(t, newName, item.GetName())
		assert.Less(t, initUpdated, item.GetUpdatedAt())
		assert.Equal(t, initCreated.Unix(), item.GetCreatedAt().Unix())
	})

	t.Run("should return error when empty name", func(t *testing.T) {
		initName := "init name"
		initCreated := time.Now()
		initUpdated := time.Now()

		item := NewItem(
			0, initName, 0, initCreated, initUpdated,
		)

		newName := ""

		item, err := domain.UpdateItemName(newName, item)

		assert.Error(t, err)
		assert.Nil(t, item)
	})

	t.Run("should return error when item is nil", func(t *testing.T) {
		newName := "new name"

		item, err := domain.UpdateItemName(newName, nil)

		assert.Error(t, err)
		assert.Nil(t, item)
	})
}
