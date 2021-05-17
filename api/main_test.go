package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeguacelestial/reto/login"
	"github.com/yeguacelestial/reto/utils"
)

type NormalResponse struct {
	Data        []map[string]string `json:"data"`
	Description string              `json:"description"`
	Message     string              `json:"message"`
}

var validUserData = User{
	Email:    "demo@usuario.com",
	Password: "pipjY7-guknaq-nancex",
}

var invalidUserData = User{
	Email:    "invalid@usuario.com",
	Password: "pipjY7-guknaq-nancex",
}

// Valid login with correct email & password
func TestValidLoginEndpoint(t *testing.T) {

	bJson, err := json.Marshal(validUserData)
	utils.HandleErr(err)

	request, err := http.NewRequest("POST", "/login", bytes.NewBuffer(bJson))
	utils.HandleErr(err)

	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "Expected 200, got another HTTP Code.")
}

// Valid login with correct email & password
func TestInvalidLoginEndpoint(t *testing.T) {
	bJson, err := json.Marshal(invalidUserData)
	utils.HandleErr(err)

	request, err := http.NewRequest("POST", "/login", bytes.NewBuffer(bJson))
	utils.HandleErr(err)

	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)

	assert.Equal(t, 400, response.Code, "Expected 400, got another HTTP Code.")
}

// Send a link, and retrieve a .xlsx file with all the links in the HTML
func TestAuthorizedGetLinksEndpoint(t *testing.T) {
	// Create a JSON body with a URL
	getLinksBody := GetLinksRequestBody{
		Url: "https://raw.githubusercontent.com/gophercises/link/master/ex1.html",
	}

	bJson, err := json.Marshal(getLinksBody)
	utils.HandleErr(err)

	request, err := http.NewRequest("POST", "/get-links", bytes.NewBuffer(bJson))
	utils.HandleErr(err)

	token, err := login.GenerateJWT(validUserData.Email, validUserData.Password)
	utils.HandleErr(err)

	bearer := "Bearer " + token

	// Add JWT
	request.Header.Add("Authorization", bearer)

	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "Expected 200, got another HTTP Code.")
}

// Calling get-links endpoint without a JWT
func TestUnauthorizedGetLinksEndpoint(t *testing.T) {
	// Create a JSON body with a URL
	getLinksBody := GetLinksRequestBody{
		Url: "https://raw.githubusercontent.com/gophercises/link/master/ex1.html",
	}

	bJson, err := json.Marshal(getLinksBody)
	utils.HandleErr(err)

	request, err := http.NewRequest("POST", "/get-links", bytes.NewBuffer(bJson))
	utils.HandleErr(err)

	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)

	assert.Equal(t, 401, response.Code, "Expected 200, got another HTTP Code.")
}

// Get the claims of a JWT token
func TestValidMeEndpoint(t *testing.T) {
	bJson, err := json.Marshal(validUserData)
	utils.HandleErr(err)

	loginRequest, err := http.NewRequest("POST", "/login", bytes.NewBuffer(bJson))
	utils.HandleErr(err)

	loginResponse := httptest.NewRecorder()

	Router().ServeHTTP(loginResponse, loginRequest)

	loginReqBody, _ := ioutil.ReadAll(loginResponse.Body)

	var normalResponse NormalResponse
	json.Unmarshal(loginReqBody, &normalResponse)

	authToken := normalResponse.Data[0]["token"]

	bearer := "Bearer " + authToken

	meRequest, err := http.NewRequest("GET", "/me", nil)
	utils.HandleErr(err)

	// Add JWT
	meRequest.Header.Add("Authorization", bearer)

	meResponse := httptest.NewRecorder()

	Router().ServeHTTP(meResponse, meRequest)

	assert.Equal(t, 200, meResponse.Code, "Expected 200, got another HTTP Code.")
}

func TestEmptyTokenMeEndpoint(t *testing.T) {
	bearer := ""

	meRequest, err := http.NewRequest("GET", "/me", nil)
	utils.HandleErr(err)

	// Add JWT
	meRequest.Header.Add("Authorization", bearer)

	meResponse := httptest.NewRecorder()

	Router().ServeHTTP(meResponse, meRequest)

	assert.Equal(t, 401, meResponse.Code, "Expected 200, got another HTTP Code.")
}

func TestRandomTokenMeEndpoint(t *testing.T) {
	bearer := "12345"

	meRequest, err := http.NewRequest("GET", "/me", nil)
	utils.HandleErr(err)

	// Add JWT
	meRequest.Header.Add("Authorization", bearer)

	meResponse := httptest.NewRecorder()

	Router().ServeHTTP(meResponse, meRequest)

	assert.Equal(t, 401, meResponse.Code, "Expected 200, got another HTTP Code.")
}
