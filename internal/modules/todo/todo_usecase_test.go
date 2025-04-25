package todo

import (
	"github.com/rykeroc/todo-cli/internal/testutils"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func beforeEach(t *testing.T) (*testutils.TestFixture, UseCase) {
	fixture := testutils.SetupTestFixture(t)
	repository := NewSqliteRepository(fixture.Db)
	domain := NewDomain()
	return fixture, NewUseCase(domain, repository)
}

func afterEach(fixture *testutils.TestFixture) {
	err := fixture.CleanupTestFixture()
	if err != nil {
		log.Fatalf("TestPersistItem_Success: Error on cleanup: %v", err)
	}
}

func TestDefaultUseCase_Create(t *testing.T) {
	fixture, useCase := beforeEach(t)
	defer afterEach(fixture)

	type testCase struct {
		name        string
		expectError bool
	}

	testCases := []testCase{
		{
			name:        "valid name",
			expectError: false,
		},
		{
			name:        "",
			expectError: true,
		},
	}

	t.Run("todo use case create", func(t *testing.T) {
		for _, test := range testCases {
			err := useCase.Create(test.name)
			if test.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		}
	})
}

func TestDefaultUseCase_List(t *testing.T) {
	fixture, useCase := beforeEach(t)
	defer afterEach(fixture)

	type testCase struct {
		expectError bool
	}

	testCases := []testCase{
		{expectError: false},
	}

	t.Run("todo use case list", func(t *testing.T) {
		for _, test := range testCases {
			err := useCase.List()
			if test.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		}
	})
}

func TestDefaultUseCase_Remove(t *testing.T) {
	fixture, useCase := beforeEach(t)
	defer afterEach(fixture)

	type testCase struct {
		itemId         int64
		expectedItemId int64
		expectError    bool
	}

	testCases := []testCase{
		{
			itemId:         1,
			expectedItemId: 1,
			expectError:    false,
		},
		{
			itemId:         100,
			expectedItemId: -1,
			expectError:    false,
		},
	}

	// Insert test item for test case 1
	err := useCase.Create("item")
	if err != nil {
		log.Fatalf("TestDefaultUseCase_Remove: Error inserting item: %v", err)
	}

	t.Run("todo use case remove", func(t *testing.T) {
		for _, test := range testCases {
			deletedId, err := useCase.Remove(test.itemId)
			if test.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.expectedItemId, deletedId)
		}
	})
}
