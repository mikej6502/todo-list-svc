package database

import "github.com/mikej6502/todo-list-svc/model"

var items = make([]model.Item, 0)

type DataStore interface {
	GetItems() []model.Item
	AddItem(item model.Item)
}

type InMemoryDataStore struct {
}

func (d InMemoryDataStore) GetItems() []model.Item {

	return items
}

func (d InMemoryDataStore) AddItem(item model.Item) {
	items = append(items, item)
}
