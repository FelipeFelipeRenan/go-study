package main

import (
	"encoding/json"
	"fmt"
	"log"
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

func addUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)
	newUser.ID = nextID
	nextID++
	users = append(users, newUser)
	json.NewEncoder(w).Encode(newUser)
}

func updateUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var updatedUser User
	json.NewDecoder(r.Body).Decode(&updatedUser)

	for i, user := range users {
		if user.ID == id{
			users[i] = updatedUser
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}
	http.Error(w, "Usuário não encontrado", http.StatusNotFound)

}

func deleteUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]... )
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "Usuario nao encontrado", http.StatusNotFound)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users", addUser).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
