package main

import (
	"html/template"
	"net/http"
)

type User struct {
	ID int
	Name string
}

var users []User

func handleRoot(w http.ResponseWriter, r *http.Request)  {
	renderTemplate(w, "index", nil)
}

func handleUsers(w http.ResponseWriter, r *http.Request)  {
	renderTemplate(w, "users", users)
}

func handleAddUsers(w http.ResponseWriter, r *http.Request)  {
	id :=  len(users) + 1
	name := r.FormValue("name")
	newUser := User{ID: id, Name: name}
	users = append(users, newUser)

	http.Redirect(w,r, "/users", http.StatusSeeOther)
}

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
}