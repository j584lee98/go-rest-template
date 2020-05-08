package main

import (
	"encoding/json"
	"net/http"
	"strconv"

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

	router.HandleFunc("/posts/{id}", getItem).Methods("GET")

	router.HandleFunc("/posts", getItems).Methods("GET")

	router.HandleFunc("/posts", addItem).Methods("POST")

	router.HandleFunc("/posts/{id}", updateItem).Methods("PUT")

	router.HandleFunc("/posts/{id}", deleteItem).Methods("DELETE")

	http.ListenAndServe(":5000", router)
}

func getItem(writer http.ResponseWriter, req *http.Request) {
	var idParam string = mux.Vars(req)["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		writer.WriteHeader(400)
		writer.Write([]byte("ID couldn't be converted into an integer"))
		return
	}

	if id >= len(items) {
		writer.WriteHeader(404)
		writer.Write([]byte("No item found with specified ID"))
		return
	}

	item := items[id]

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(item)
}

func getItems(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(items)
}

func addItem(writer http.ResponseWriter, req *http.Request) {
	var newItem Item
	json.NewDecoder(req.Body).Decode(&newItem)

	items = append(items, newItem)

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(items)
}

func updateItem(writer http.ResponseWriter, req *http.Request) {
	var idParam string = mux.Vars(req)["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		writer.WriteHeader(400)
		writer.Write([]byte("ID couldn't be converted into an integer"))
		return
	}

	if id >= len(items) {
		writer.WriteHeader(404)
		writer.Write([]byte("No item found with specified ID"))
		return
	}

	var updatedItem Item
	json.NewDecoder(req.Body).Decode(&updatedItem)

	items[id] = updatedItem
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(updatedItem)
}

func deleteItem(writer http.ResponseWriter, req *http.Request) {
	var idParam string = mux.Vars(req)["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		writer.WriteHeader(400)
		writer.Write([]byte("ID couldn't be converted into an integer"))
		return
	}

	if id >= len(items) {
		writer.WriteHeader(404)
		writer.Write([]byte("No item found with specified ID"))
		return
	}

	items = append(items[:id], items[id+1:]...)

	writer.WriteHeader(200)
}
