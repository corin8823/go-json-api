package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func decorator(h func(r *http.Request) Response) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result := h(r)
		result.Write(w)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", decorator(Logging(Index, "index"))).Methods("GET")
	r.HandleFunc("/todos", decorator(Logging(TodoIndex, "todo-index"))).Methods("GET")
	r.HandleFunc("/todos/{todoId}", decorator(IDShouldBeInt(TodoShow, "todo-show"))).Methods("GET")
	r.HandleFunc("/todos", decorator(Logging(TodoCreate, "todo-create"))).Methods("POST")
	r.HandleFunc("/todos/{todoId}", decorator(IDShouldBeInt(TodoDelete, "todo-delete"))).Methods("DELETE")

	http.Handle("/", r)

	log.Println("start")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
