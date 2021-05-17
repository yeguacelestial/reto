package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/yeguacelestial/reto/getlinks"
	"github.com/yeguacelestial/reto/login"
	"github.com/yeguacelestial/reto/utils"

	"github.com/gorilla/mux"
)

type Dictionary map[string]interface{}

// Email and password default structure.
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Global Users slice. Simulates a database.
var Users []User

// 'get-links' endpoint should receive an url and a bearer token
// for  a valid response.
type GetLinksRequestBody struct {
	Url string `json:"Url"`
}

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

	handleRequests(router)

	return router
}

// Handle each request from router
func handleRequests(router *mux.Router) {
	router.HandleFunc("/login", LoginEndpoint).Methods("POST")
	router.Handle("/get-links", login.IsAuthorized(GetLinksEndpoint)).Methods("POST")
	router.Handle("/me", login.IsAuthorized(MeEndpoint)).Methods("GET")
}

func LoginEndpoint(response http.ResponseWriter, request *http.Request) {
	// get the body of our POST request
	reqBody, _ := ioutil.ReadAll(request.Body)

	// Unmarshal this into new User struct
	var user User
	json.Unmarshal(reqBody, &user)

	k, found := FindUser(Users, user)

	jsonData := simplejson.New()

	// If user was not found in database...
	if !found && k == -1 {
		data := []Dictionary{
			{
				"email":    user.Email,
				"password": user.Password,
			},
		}

		// Set the JSON Body values
		response.WriteHeader(400)
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

		validToken, err := login.GenerateJWT(user.Email, user.Password)
		if err != nil {
			fmt.Println("[-] Error generating JWT => ", err.Error())
		}

		data := []Dictionary{
			{
				"email":    user.Email,
				"password": user.Password,
				"token":    validToken,
			},
		}

		// Set the JSON Body values
		response.WriteHeader(200)
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

	var excelRows [][]map[string]string
	excelRows = utils.ArrayForExcel(excelRows, "TEXT", "HREF")

	// 4. Iterate on the slice of Link struct
	for i := 0; i < len(htmlLinks); i++ {
		tagText := htmlLinks[i].Text
		tagHref := htmlLinks[i].Href

		excelRows = utils.ArrayForExcel(excelRows, tagText, tagHref)
	}

	// 5. Convert struct of links and texts to a xlsx file
	f := utils.CreateSheet(nil, "Challenge", excelRows)
	utils.CreateExcel(f, "extractedLinks.xlsx")

	// 6. Add file to response
	response.Header().Set("Content-Type", "application/octet-stream")
	response.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote("extractedLinks.xlsx"))

	// 7. Serve file
	http.ServeFile(response, request, "extractedLinks.xlsx")

	// Remove file
	e := os.Remove("extractedLinks.xlsx")
	if e != nil {
		log.Fatal(e)
	}
}

func MeEndpoint(response http.ResponseWriter, request *http.Request) {
	// Get Authorization header
	authorizationHeader := request.Header.Get("Authorization")

	// Extract token
	splitToken := strings.Split(authorizationHeader, "Bearer ")
	reqToken := splitToken[1]

	claims, ok := login.ExtractClaims(reqToken)
	if !ok {
		log.Fatal(ok)
	}

	// Set the JSON Body values
	jsonData := simplejson.New()

	response.Header().Set("Content-Type", "application/json")
	jsonData.Set("message", "success")
	jsonData.Set("description", "fetched user data from claims")
	jsonData.Set("data", claims)

	payload, err := jsonData.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}

	response.Write(payload)
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
