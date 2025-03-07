package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// User struct
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

func main() {
	router := GetRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	users = append(users, newUser)
	err := json.NewEncoder(w).Encode(newUser)
	if err != nil {
		log.Fatal(err)
		return
	}
}
