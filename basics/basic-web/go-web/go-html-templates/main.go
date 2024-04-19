package main

import (
	"html/template"
	"net/http"
)

// struct to represent the user data
type User struct {
	ID int
	Name string
}
// user list
var users []User

// functions to handle wich file will be renderized
func handleRoot(w http.ResponseWriter, r *http.Request)  {
	renderTemplate(w, "index", nil)
}

func handleUsers(w http.ResponseWriter, r *http.Request)  {
	renderTemplate(w, "users", users)
}

// function to handle the user adding
func handleAddUsers(w http.ResponseWriter, r *http.Request)  {
	id :=  len(users) + 1
	name := r.FormValue("name")
	newUser := User{ID: id, Name: name}
	users = append(users, newUser)

	http.Redirect(w,r, "/users", http.StatusSeeOther)
}

// function to render the html files from templates folder
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}){
	
	t, err := template.ParseFiles("templates/" + tmpl + ".html")
	
	if err != nil {
		http.Error(w, "Erro ao carregar o template", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
	}
}  
	


func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/users", handleUsers)
	http.HandleFunc("/users/add", handleAddUsers)

	http.ListenAndServe(":8080", nil)
}