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
	if r.Method == http.MethodGet {
		p := strings.Split(r.URL.Path, "/")[1:]
		n := len(p)

		if p[0] == "todo" && p[1] == "" {
			listItems(w, r)
		} else if n == 2 && p[0] == "todo" {
			getItemById(w, r, p[1])
		}
	} else if r.Method == http.MethodPost {
		addItem(w, r)
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
		var json, _ = json.Marshal(dataStore.AddItem(item))
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)

		w.Write(json)
	}
}
