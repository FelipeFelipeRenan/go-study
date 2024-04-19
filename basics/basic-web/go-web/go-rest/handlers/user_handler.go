package handlers

import (
    "encoding/json"
    "main/models"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
    users := models.GetAllUsers()
    json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    user, err := models.GetUserByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(user)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
    var newUser models.User
    json.NewDecoder(r.Body).Decode(&newUser)
    addedUser := models.AddUser(newUser)
    json.NewEncoder(w).Encode(addedUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    var updatedUser models.User
    json.NewDecoder(r.Body).Decode(&updatedUser)
    _, err := models.UpdateUser(updatedUser)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(updatedUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    deletedUser, err := models.DeleteUser(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(deletedUser)
}
