package models

import "fmt"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User
var nextID = 1

func GetAllUsers() []User {
	return users
}

func GetUserByID(id int) (*User, error) {
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("Usuário não encontrado")
}

func AddUser(newUser User) User {
	newUser.ID = nextID
	nextID++
	users = append(users, newUser)
	return newUser
}

func UpdateUser(updatedUser User) (*User, error) {
	for i, user := range users {
		if user.ID == updatedUser.ID {
			users[i] = updatedUser
			return &updatedUser, nil
		}
	}
	return nil, fmt.Errorf("Usuário não encontrado")
}

func DeleteUser(id int) (*User, error) {
	for i, user := range users {
		if user.ID == id {
			deletedUser := users[i]
			users = append(users[:i], users[i+1:]...)
			return &deletedUser, nil
		}
	}
	return nil, fmt.Errorf("Usuário não encontrado")
}
