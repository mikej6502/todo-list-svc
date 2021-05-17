package controllers

import (
	"encoding/json"
	"github.com/mikej6502/todo-list-svc/database"
	"github.com/mikej6502/todo-list-svc/model"
	"net/http"
)

var inMemoryDataStore = database.InMemoryDataStore{}

func ProcessRequest(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		listItems(w, r)
	} else if r.Method == http.MethodPost {
		addItems(w, r)
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

// GET all items in the to do list
func listItems(w http.ResponseWriter, r *http.Request) {
	var json, _ = json.Marshal(inMemoryDataStore.GetItems())

	w.Write(json)
}

// POST a new item to the to do list
func addItems(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var item model.Item
	err := decoder.Decode(&item)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		inMemoryDataStore.AddItem(item)
		w.WriteHeader(http.StatusCreated)
	}
}
