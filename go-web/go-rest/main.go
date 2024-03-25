package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "main/handlers"
)

func main() {
    router := mux.NewRouter()

    // Rotas para usuários
    router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
    router.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
    router.HandleFunc("/users", handlers.AddUser).Methods("POST")
    router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
    router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

    fmt.Println("Servidor rodando em http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
