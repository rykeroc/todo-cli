package todo

import "fmt"

type UseCase interface {
	Create(string) error
	List() error
}

type defaultUseCase struct {
	domain     Domain
	repository Repository
}

func NewUseCase(repository Repository) UseCase {
	return &defaultUseCase{
		domain:     NewDomain(),
		repository: repository,
	}
}

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

func (uc *defaultUseCase) List() error {
	items, err := uc.repository.FindAllItems()
	if err != nil {
		return fmt.Errorf("defaultUseCase.List: %v", err)
	}
	tabularList, err := uc.domain.GetItemList(items)
	if err != nil {
		return fmt.Errorf("defaultUseCase.List: %v", err)
	}

	fmt.Println(tabularList)
	return nil
}
