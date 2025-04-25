package todo

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewItem(t *testing.T) {
	var id int64 = 0
	name := "name"
	var isCompleted int8 = 0
	updatedAt := time.Now()
	createdAt := updatedAt

	item := NewItem(id, name, isCompleted, updatedAt, createdAt)

	assert.Equal(t, id, item.GetId())
	assert.Equal(t, name, item.GetName())
	assert.Equal(t, isCompleted, item.GetIsCompleted())
	assert.Equal(t, updatedAt, item.GetUpdatedAt())
	assert.Equal(t, createdAt, item.GetCreatedAt())
}
