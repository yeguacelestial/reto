package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

// Global Users slice. Simulates a database.
var Users []User

func main() {

	// Create User
	defaultUser := User{
		Email:    "demo@usuario.com",
		Password: "pipjY7-guknaq-nancex",
	}

	Users = append(Users, defaultUser)

	router := Router()
	log.Fatal(http.ListenAndServe(":10000", router))
}

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", RootEndpoint).Methods("GET")
	router.HandleFunc("/login", LoginEndpoint).Methods("POST")

	return router
}

func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
	response.Write([]byte("Hello World"))
}

func LoginEndpoint(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	reqBody, _ := ioutil.ReadAll(r.Body)

	// Unmarshal this into new User struct
	var user User
	json.Unmarshal(reqBody, &user)

	k, found := Find(Users, user)

	if !found {
		fmt.Printf("[-] Invalid User or password.\n")
	} else {
		fmt.Printf("[+] User with email %s is registered in the db. (Index: %d)\n", Users[k].Email, k)
	}
}

// Verify if user is registered in the database (slice)
func Find(slice []User, val User) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}

	return -1, false
}
