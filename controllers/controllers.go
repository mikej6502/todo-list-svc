package controllers

import (
	"encoding/json"
	"github.com/mikej6502/todo-list-svc/database"
	"github.com/mikej6502/todo-list-svc/model"
	"net/http"
)

// GET all items in the to do list
func ProcessRequest(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		listItems(w, r)
	} else if r.Method == http.MethodPost {
		addItems(w, r)
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func listItems(w http.ResponseWriter, r *http.Request) {
	datastore := database.InMemoryDataStore{}

	var json, _ = json.Marshal(datastore.GetItems())

	w.Write(json)
}

func addItems(w http.ResponseWriter, r *http.Request) {
	datastore := database.InMemoryDataStore{}

	decoder := json.NewDecoder(r.Body)
	var item model.Item
	err := decoder.Decode(&item)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		datastore.AddItem(item)
		w.WriteHeader(http.StatusCreated)
	}
}
