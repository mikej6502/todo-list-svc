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

func (d InMemoryDataStore) AddItem(item model.Item) model.Item {
	item.Id = strconv.FormatInt(nextId, 10)
	nextId++

	items = append(items, item)
	return item
}
