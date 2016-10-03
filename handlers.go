package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome")
}

func TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todos := []Todo{ Todo{Name: "Write presentation"}, Todo{Name: "Host meetup"} }

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "TODO show: %s", ps.ByName("todoId"))
}
