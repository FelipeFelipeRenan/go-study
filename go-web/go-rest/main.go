package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User
var nextID = 1

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, user := range users {
		if strconv.Itoa(user.ID) == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}

	}

	http.Error(w, "Usuario não encontrado", http.StatusNotFound)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users")
}
