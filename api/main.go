package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/yeguacelestial/reto/getlinks"
	"github.com/yeguacelestial/reto/utils"

	"github.com/gorilla/mux"
)

type Dictionary map[string]interface{}

// Email and password default structure.
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// 'get-links' endpoint should receive an url and a bearer token
// for  a valid response.
type GetLinksRequestBody struct {
	Url string `json:"Url"`
}

// Global Users slice. Simulates a database.
var Users []User

// Server port
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

	// Handle endpoints
	router.HandleFunc("/login", LoginEndpoint).Methods("POST")
	router.HandleFunc("/get-links", GetLinksEndpoint).Methods("POST")

	return router
}

func LoginEndpoint(response http.ResponseWriter, request *http.Request) {
	// get the body of our POST request
	reqBody, _ := ioutil.ReadAll(request.Body)

	// Unmarshal this into new User struct
	var user User
	json.Unmarshal(reqBody, &user)

	k, found := FindUser(Users, user)

	jsonData := simplejson.New()

	if !found && k == -1 {

		response.WriteHeader(401)

		data := []Dictionary{
			{
				"email": user.Email,
			},
		}

		// Set the JSON Body values
		response.Header().Set("Content-Type", "application/json")
		jsonData.Set("message", "error")
		jsonData.Set("description", "invalid email or password")
		jsonData.Set("data", data)

		payload, err := jsonData.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}

		response.Write(payload)

		// Else, returns an error indicating that the user is valid.
	} else {
		response.WriteHeader(200)

		data := []Dictionary{
			{
				"email": user.Email,
			},
		}

		// Set the JSON Body values
		response.Header().Set("Content-Type", "application/json")
		jsonData.Set("message", "success")
		jsonData.Set("description", "logged in successfully")
		jsonData.Set("data", data)

		payload, err := jsonData.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}

		response.Write(payload)
	}
}

func GetLinksEndpoint(response http.ResponseWriter, request *http.Request) {
	// 1. Read the url from the response
	reqBody, _ := ioutil.ReadAll(request.Body)

	// Unmarshal this into new User struct
	var link GetLinksRequestBody
	json.Unmarshal(reqBody, &link)

	url := link.Url

	// 2. Parse html from the url string
	htmlString := utils.ParseHtmlFromUrl(url)
	htmlStringReader := strings.NewReader(htmlString)

	// 3. Extract all the links from the html
	htmlLinks, err := getlinks.ParseLinksFromHtmlReader(htmlStringReader)
	utils.HandleErr(err)

	// 4. Convert slice of links and texts to a csv file
	fmt.Println(htmlLinks)

	// 5. Add file to response
}

// Verify if user is registered in the 'database' (users slice)
func FindUser(users []User, val User) (int, bool) {
	for i, item := range users {
		if item == val {
			return i, true
		}
	}

	return -1, false
}
