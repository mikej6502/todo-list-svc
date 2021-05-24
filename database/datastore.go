package database

import "github.com/mikej6502/todo-list-svc/model"

type DataStore interface {
	GetItem(id string) (model.Item, error)
	GetItems() []model.Item
	AddItem(item model.Item) (model.Item, error)
	UpdateItem(item model.Item, id string) error
}
