package main

import (
	"github.com/mikej6502/todo-list-svc/controllers"
	"github.com/mikej6502/todo-list-svc/database"
	"log"
	"net/http"
)

func main() {
	//var dataStore = database.InMemoryDataStore{}
	var dataStore = database.MongoDBDataStore{Url: "mongodb://localhost:27017/",
		DatabaseName:   "todo-list-db",
		CollectionName: "todo-list"}

	controllers.Initialise(&dataStore)

	http.HandleFunc("/todo", controllers.ProcessRequest)
	http.HandleFunc("/todo/", controllers.ProcessRequest)

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", http.StripPrefix("/", fs))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
