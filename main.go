package main

import (
	"github.com/mikej6502/todo-list-svc/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/todo", controllers.ProcessRequest)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
