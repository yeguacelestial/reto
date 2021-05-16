package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Valid login with correct email & password
func TestValidLoginEndpoint(t *testing.T) {
	userData := User{
		Email:    "demo@usuario.com",
		Password: "pipjY7-guknaq-nancex",
	}

	bJson, err := json.Marshal(userData)
	handleErr(err)

	request, err := http.NewRequest("POST", "/login", bytes.NewBuffer(bJson))
	handleErr(err)

	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)

	expectedResponseBody := `{"data":[{"email":"demo@usuario.com"}],"description":"logged in successfully","message":"success"}`

	assert.Equal(t, 200, response.Code, "Expected 200, got another HTTP Code.")
	assert.Equal(t, expectedResponseBody, response.Body.String(), "Unexpected response body.")
}

// Valid login with correct email & password
func TestInvalidLoginEndpoint(t *testing.T) {
	userData := User{
		Email:    "invalid@usuario.com",
		Password: "pipjY7-guknaq-nancex",
	}

	bJson, err := json.Marshal(userData)
	handleErr(err)

	request, err := http.NewRequest("POST", "/login", bytes.NewBuffer(bJson))
	handleErr(err)

	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)

	expectedResponseBody := `{"data":[{"email":"invalid@usuario.com"}],"description":"invalid email or password","message":"error"}`

	assert.Equal(t, 401, response.Code, "Expected 401, got another HTTP Code.")
	assert.Equal(t, expectedResponseBody, response.Body.String(), "Unexpected response body.")
}

// Send a link, and retrieve a .csv file with all the links in the HTML
func TestGetLinksEndpoint(t *testing.T) {
	// Create a JSON body with a URL
	getLinksBody := Link{
		Url: "https://raw.githubusercontent.com/gophercises/link/master/ex1.html",
	}

	bJson, err := json.Marshal(getLinksBody)
	handleErr(err)

	request, err := http.NewRequest("POST", "/get-links", bytes.NewBuffer(bJson))
	handleErr(err)

	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)

	expectedResponseBody := `{"data":[{"file":"path_to_file.csv"}],"description":"generated csv file","message":"success"}`

	assert.Equal(t, expectedResponseBody, response.Body.String(), "Unexpected JSON Body.")
}
