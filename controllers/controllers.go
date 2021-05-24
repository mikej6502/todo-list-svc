package controllers

import (
	"encoding/json"
	"github.com/mikej6502/todo-list-svc/database"
	"github.com/mikej6502/todo-list-svc/model"
	"net/http"
	"strings"
)

var dataStore = database.InMemoryDataStore{}

func ProcessRequest(w http.ResponseWriter, r *http.Request) {
	p := strings.Split(r.URL.Path, "/")[1:]
	n := len(p)

	if r.Method == http.MethodGet {
		if p[0] == "todo" && p[1] == "" {
			listItems(w, r)
		} else if n == 2 && p[0] == "todo" {
			getItemById(w, r, p[1])
		}
	} else if r.Method == http.MethodPost {
		addItem(w, r)
	} else if r.Method == http.MethodPut {
		if len(p) == 2 && p[0] == "todo" {
			updateItem(w, r, p[1])
		}
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

// GET all items in the to do list
func getItemById(w http.ResponseWriter, r *http.Request, id string) {

	var item, err = dataStore.GetItem(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	} else {
		var json, _ = json.Marshal(item)
		w.Header().Add("content-type", "application/json")
		w.Write(json)
	}
}

// GET single item by id
func listItems(w http.ResponseWriter, r *http.Request) {
	var json, _ = json.Marshal(dataStore.GetItems())

	w.Header().Add("content-type", "application/json")
	w.Write(json)
}

// POST a new item to the to do list
func addItem(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var item model.Item
	err := decoder.Decode(&item)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		var newItem, err = dataStore.AddItem(item)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		json, _ := json.Marshal(newItem)

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(json)
	}
}

// PUT: Update an existing item in the to do list
func updateItem(w http.ResponseWriter, r *http.Request, id string) {
	decoder := json.NewDecoder(r.Body)
	var item model.Item
	err := decoder.Decode(&item)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		err := dataStore.UpdateItem(item, id)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}
