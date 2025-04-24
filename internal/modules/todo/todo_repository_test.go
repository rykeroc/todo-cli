package todo

import (
	"github.com/rykeroc/todo-cli/internal/testutils"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var testItems = []Item{
	NewItem(
		0,
		"item",
		time.Now(),
		time.Now(),
	),
	NewItem(
		0,
		"another item",
		time.Now(),
		time.Now(),
	),
}

func TestPersistItem_Success(t *testing.T) {
	fixture := testutils.SetupTestFixture(t)
	defer func(fixture *testutils.TestFixture) {
		err := fixture.CleanupTestFixture()
		if err != nil {
			log.Fatalf("TestPersistItem_Success: Error on cleanup: %v", err)
		}
	}(fixture)
	repository := NewSqliteRepository(fixture.Db)

	t.Run("should persist item successfully", func(t *testing.T) {
		for _, item := range testItems {
			insertedId, err := repository.PersistItem(item)
			assert.ErrorIs(t, nil, err)
			assert.Greater(t, insertedId, int64(0))
		}
	})
}

func TestPersistItem_NoDatabase(t *testing.T) {
	repository := NewSqliteRepository(nil)

	t.Run("should have error on no database", func(t *testing.T) {
		for _, item := range testItems {
			id, err := repository.PersistItem(item)
			assert.Error(t, err)
			assert.Equal(t, int64(-1), id)
		}
	})
}

func TestFindAllItems_Success(t *testing.T) {
	fixture := testutils.SetupTestFixture(t)
	defer func(fixture *testutils.TestFixture) {
		err := fixture.CleanupTestFixture()
		if err != nil {
			log.Fatalf("TestFindAllItems_Success: Error on cleanup: %v", err)
		}
	}(fixture)
	repository := NewSqliteRepository(fixture.Db)

	t.Run("should get all items successfully", func(t *testing.T) {
		for _, item := range testItems {
			_, err := repository.PersistItem(item)
			if err != nil {
				t.Fatalf("TestFindAllItems_Success: %v", err)
			}
		}

		result, err := repository.FindAllItems()
		assert.NoError(t, err)
		assert.Equal(t, len(result), 2)
	})
}

func TestFindAllItems_NoDatabase(t *testing.T) {
	repository := NewSqliteRepository(nil)

	t.Run("should return error on no database", func(t *testing.T) {
		items, err := repository.FindAllItems()
		assert.Error(t, err)
		assert.Equal(t, 0, len(items))
	})
}

func TestUpdateItemById_Success(t *testing.T) {
	fixture := testutils.SetupTestFixture(t)
	defer func(fixture *testutils.TestFixture) {
		err := fixture.CleanupTestFixture()
		if err != nil {
			log.Fatalf("TestSqliteRepository_UpdateItemById_Success: Error on cleanup: %v", err)
		}
	}(fixture)
	repository := NewSqliteRepository(fixture.Db)

	t.Run("should update item successfully", func(t *testing.T) {
		_, err := repository.PersistItem(testItems[0])
		if err != nil {
			t.Fatalf("TestUpdateItemById_Success: %v", err)
		}

		itemToUpdate := NewItem(
			1,
			"new name",
			time.Now(),
			testItems[0].GetCreatedAt(),
		)

		affectedRows, err := repository.UpdateItemById(itemToUpdate)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), affectedRows)
	})
}

func TestUpdateItemById_ItemNotExist(t *testing.T) {
	fixture := testutils.SetupTestFixture(t)
	defer func(fixture *testutils.TestFixture) {
		err := fixture.CleanupTestFixture()
		if err != nil {
			log.Fatalf("TestSqliteRepository_UpdateItemById_Success: Error on cleanup: %v", err)
		}
	}(fixture)
	repository := NewSqliteRepository(fixture.Db)

	t.Run("should affect 0 rows", func(t *testing.T) {
		itemToUpdate := testItems[0]
		itemToUpdate.SetName("new name")
		itemToUpdate.SetUpdatedAt(time.Now())

		affectedRows, err := repository.UpdateItemById(itemToUpdate)
		assert.NoError(t, err)
		assert.Equal(t, int64(0), affectedRows)
	})
}

func TestUpdateItemById_NoDatabase(t *testing.T) {
	repository := NewSqliteRepository(nil)

	t.Run("should return error on no database", func(t *testing.T) {
		result, err := repository.UpdateItemById(testItems[0])
		assert.Error(t, err)
		assert.Equal(t, int64(-1), result)
	})
}

func TestDeleteItem_Success(t *testing.T) {
	fixture := testutils.SetupTestFixture(t)
	defer func(fixture *testutils.TestFixture) {
		err := fixture.CleanupTestFixture()
		if err != nil {
			log.Fatalf("TestSqliteRepository_PersistItem: Error on cleanup: %v", err)
		}
	}(fixture)
	repository := NewSqliteRepository(fixture.Db)

	t.Run("should delete existing todo", func(t *testing.T) {
		_, err := repository.PersistItem(testItems[0])
		if err != nil {
			t.Fatalf("TestSqliteRepository_UpdateItemById_Success: %v", err)
		}

		affectedRows, err := repository.DeleteItemById(1)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), affectedRows)
	})
}

func TestDeleteItem_ItemNotExists(t *testing.T) {
	fixture := testutils.SetupTestFixture(t)
	defer func(fixture *testutils.TestFixture) {
		err := fixture.CleanupTestFixture()
		if err != nil {
			log.Fatalf("TestSqliteRepository_PersistItem: Error on cleanup: %v", err)
		}
	}(fixture)
	repository := NewSqliteRepository(fixture.Db)

	t.Run("should affect 0 rows", func(t *testing.T) {
		affectedRows, err := repository.UpdateItemById(testItems[0])
		assert.NoError(t, err)
		assert.Equal(t, int64(0), affectedRows)
	})
}

func TestDeleteItem_NoDatabase(t *testing.T) {
	repository := NewSqliteRepository(nil)

	t.Run("should return error on no database", func(t *testing.T) {
		result, err := repository.DeleteItemById(1)
		assert.Error(t, err)
		assert.Equal(t, int64(-1), result)
	})
}
