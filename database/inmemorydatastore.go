package database

import (
	"errors"
	"github.com/mikej6502/todo-list-svc/model"
	"strconv"
)

var items = make([]model.Item, 0)

var nextId int64

type InMemoryDataStore struct {
}

func (d InMemoryDataStore) GetItem(id string) (model.Item, error) {
	for _, item := range items {
		if id == item.Id {
			return item, nil
		}
	}

	return model.Item{}, errors.New("item not found")
}

func (d InMemoryDataStore) GetItems() []model.Item {
	return items
}

func (d InMemoryDataStore) AddItem(item model.Item) (model.Item, error) {
	item.Id = strconv.FormatInt(nextId, 10)
	nextId++

	items = append(items, item)

	return item, nil
}

func (d InMemoryDataStore) UpdateItem(item model.Item, id string) error {
	var updated = false

	for i := range items {
		var retrievedItem = &items[i]
		if retrievedItem.Id == id {
			retrievedItem.Title = item.Title
			retrievedItem.Description = item.Description
			retrievedItem.Completed = item.Completed

			updated = true
			break
		}
	}

	if !updated {
		return errors.New("cannot find item for id: " + id)
	}

	return nil
}
