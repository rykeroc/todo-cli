package todo

import "fmt"

// UseCase godoc
//
// An interface that defines the behaviour for a todo item use case struct.
type UseCase interface {
	Create(string) error
	List() error
	Remove(int64) (int64, error)
	Update(int64, string) (int64, error)
}

// defaultUseCase godoc
//
// A structure which takes a todo domain and repository.
//
// Adheres to the todo UseCase interface.
type defaultUseCase struct {
	domain     Domain
	repository Repository
}

// NewUseCase godoc
//
// Creates a new UseCase with the passed in Domain and Repository instances.
func NewUseCase(domain Domain, repository Repository) UseCase {
	return &defaultUseCase{
		domain:     domain,
		repository: repository,
	}
}

// Create godoc
//
// Construct a new todo item using the passed in name and persist it locally.
//
// Returns error on error, nil otherwise.
func (uc *defaultUseCase) Create(name string) error {
	item, err := uc.domain.CreateItem(name)
	if err != nil {
		return fmt.Errorf("defaultUseCase.Create: %v", err)
	}
	_, err = uc.repository.PersistItem(item)
	if err != nil {
		return err
	}
	fmt.Printf("Created new todo: %s\n", name)
	return nil
}

// List godoc
//
// Get all persisted todo items and print them in a tabular list.
//
// Returns error on error, nil otherwise.
func (uc *defaultUseCase) List() error {
	items, err := uc.repository.FindAllItems()
	if err != nil {
		return fmt.Errorf("defaultUseCase.List: %v", err)
	}
	tabularList, err := uc.domain.GetTabularItemList(items)
	if err != nil {
		return fmt.Errorf("defaultUseCase.List: %v", err)
	}

	fmt.Println(tabularList)
	return nil
}

// Remove godoc
//
// Remove an item by its ID.
//
// Returns -1 and nil when the item does not exist.
//
// Returns -1 and error on error.
//
// Returns the deleted item's ID when the item is deleted successfully.
func (uc *defaultUseCase) Remove(itemId int64) (int64, error) {
	// Invalid item ID
	if itemId == 0 {
		return -1, nil
	}

	// Delete the item by its ID
	affectedRows, err := uc.repository.DeleteItemById(itemId)
	if err != nil {
		return -1, fmt.Errorf("defaultUseCase.Remove: Failed to delete item with ID %d: %v", itemId, err)
	}
	// Item not deleted as it does not exist
	if affectedRows == 0 {
		return -1, nil
	}

	return itemId, nil
}

// Update godoc
//
// Update an item's name by itemId.
//
// Returns -1 and nil if the item does not exist.
//
// Returns -1 and error on error.
//
// Returns updated item id and nil on success.
func (uc *defaultUseCase) Update(itemId int64, newName string) (int64, error) {
	// Invalid item ID
	if itemId == 0 {
		return -1, nil
	}
	// New name is empty
	if len(newName) == 0 {
		return -1, fmt.Errorf("defaultUseCase.Update: New name for item is empty")
	}

	// Find item by ID
	foundItem, err := uc.repository.FindItemById(itemId)
	// Error occurred while finding item
	if err != nil {
		return -1, fmt.Errorf("defaultUseCase.Update: Failed to find item with ID %d: %v", itemId, err)
	}
	// Item not found
	if foundItem == nil {
		return -1, nil
	}

	// Update item
	updatedItem, err := uc.domain.UpdateItem(newName, foundItem)
	// Error occurred while updating item
	if err != nil {
		return -1, fmt.Errorf("defaultUseCase.Update: Failed to update item with ID %d: %v", itemId, err)
	}

	// Delete the item by its ID
	affectedRows, err := uc.repository.UpdateItemById(updatedItem)
	// Error while persisting item update
	if err != nil {
		return -1, fmt.Errorf("defaultUseCase.Update: Failed to persists update for item with ID %d: %v", itemId, err)
	}
	// Item not deleted as it does not exist
	if affectedRows == 0 {
		return -1, nil
	}

	return itemId, nil
}
