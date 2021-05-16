package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"

	"github.com/gorilla/mux"
)

type User struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type ResponseContent struct {
	Message     string `json:"Message"`
	Description string `json:"Description"`
	Data        string `json:"Data"`
}

// Global Users slice. Simulates a database.
var Users []User

var port string = ":10000"

func main() {
	fmt.Println("[*] REST API - Mux Router")
	fmt.Println("[*] Serving on port " + port + "\n")
	router := Router()
	log.Fatal(http.ListenAndServe(port, router))
}

func Router() *mux.Router {
	// Create default User
	defaultUser := User{
		Email:    "demo@usuario.com",
		Password: "pipjY7-guknaq-nancex",
	}

	Users = append(Users, defaultUser)

	fmt.Println("[*] Created default user on database with email: " + defaultUser.Email)

	// Init router
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", RootEndpoint).Methods("GET")
	router.HandleFunc("/login", LoginEndpoint).Methods("POST")

	return router
}

func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
	response.Write([]byte("Hello World"))
}

func LoginEndpoint(response http.ResponseWriter, request *http.Request) {
	// get the body of our POST request
	reqBody, _ := ioutil.ReadAll(request.Body)

	// Unmarshal this into new User struct
	var user User
	json.Unmarshal(reqBody, &user)

	k, found := FindUser(Users, user)

	jsonData := simplejson.New()

	// If user is validated, returns a JWT response with the email, and the secret word.
	if !found {

		response.WriteHeader(401)

		// Set the JSON Body values
		response.Header().Set("Content-Type", "application/json")
		jsonData.Set("message", "error")
		jsonData.Set("description", "invalid email or password")
		jsonData.Set("email", user.Email)
		jsonData.Set("password", user.Password)

		payload, err := jsonData.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}

		response.Write(payload)

		// Else, returns an error indicating that the user is not valid.
	} else {
		response.WriteHeader(200)

		// Set the JSON Body values
		response.Header().Set("Content-Type", "application/json")
		jsonData.Set("message", "success")
		jsonData.Set("description", "logged in successfully")
		jsonData.Set("email", Users[k].Email)

		payload, err := jsonData.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}

		response.Write(payload)
	}
}

// Verify if user is registered in the 'database' (slice)
func FindUser(slice []User, val User) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}

	return -1, false
}
