package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var items = []Item{
	{ID: 1, Name: "Laptop"},
	{ID: 2, Name: "Phone"},
}

func getItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}

func getItemByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/items/")
	id, _ := strconv.Atoi(idStr)

	for _, item := range items {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	json.NewDecoder(r.Body).Decode(&newItem)
	newItem.ID = len(items) + 1
	items = append(items, newItem)
	json.NewEncoder(w).Encode(newItem)
}

func main() {
	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getItems(w, r)
		} else if r.Method == http.MethodPost {
			createItem(w, r)
		}
	})

	http.HandleFunc("/items/", getItemByID)

	http.ListenAndServe(":8081", nil)
}
