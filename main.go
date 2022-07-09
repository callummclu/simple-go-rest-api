package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/users", returnAllUsers)
	myRouter.HandleFunc("/user", createNewUser).Methods("POST")

	myRouter.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{id}", returnSingleUser)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllUsers")
	json.NewEncoder(w).Encode(Users)
}

func returnSingleUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, user := range Users {
		if user.Id == key {
			json.NewEncoder(w).Encode(user)
		}
	}
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)
	Users = append(Users, user)

	json.NewEncoder(w).Encode(user)

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, user := range Users {
		if user.Id == id {
			Users = append(Users[:index], Users[index+1:]...)
		}
	}
}

type User struct {
	Id       string `json:"Id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var Users []User

func main() {
	fmt.Println("REST API Test Project with Mux Router")
	Users = []User{
		{Id: "1", Name: "callum", Username: "callummclu", Email: "callummcluskey@mclu.co.uk"},
		{Id: "2", Name: "john", Username: "johnappleased", Email: "johnappleased@apple.com"},
	}

	handleRequests()
}
