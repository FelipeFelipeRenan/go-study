package main

import (
	"fmt"
	"net/http"
)

// functions to deal with the routes of requests
func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bem vindo a pagina inicial!")

}

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Olá, mundo!")
}

func main() {

	// assingment of the routes to the handlers 
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/hello", handleHello)

	fmt.Println("Servidor rodando em http://localhost:8080")

	// server up and running 
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
	}
}
