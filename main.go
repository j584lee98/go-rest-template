package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// User struct
type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Item struct
type Item struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	User User   `json:"user"`
}

var items []Item = []Item{}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/add", addItem).Methods("POST")

	http.ListenAndServe(":5000", router)
}

func addItem(writer http.ResponseWriter, req *http.Request) {
	var newItem Item
	json.NewDecoder(req.Body).Decode(&newItem)

	items = append(items, newItem)

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(items)
}
