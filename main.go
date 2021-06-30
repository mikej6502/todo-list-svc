package main

import (
	"github.com/mikej6502/todo-list-svc/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/todo", controllers.ProcessRequest)
	http.HandleFunc("/todo/", controllers.ProcessRequest)

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", http.StripPrefix("/", fs))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
