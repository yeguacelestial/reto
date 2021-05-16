package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func TestRootEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "OK Response is expected")
	assert.Equal(t, "Hello World", response.Body.String(), "Incorrect body found")
}

// Login with default email & password
func TestLoginEndpoint(t *testing.T) {
	jsonStr := User{
		Email:    "demo@usuario.com",
		Password: "pipjY7-guknaq-nancex",
	}

	bJson, err := json.Marshal(jsonStr)
	handleErr(err)

	request, err := http.NewRequest("POST", "/login", bytes.NewBuffer(bJson))
	handleErr(err)

	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)

	fmt.Print(response.Body.String())

	assert.Equal(t, 200, response.Code, "Expected 200, got another HTTP Code.")
}

// Send a link, and retrieve a .csv file with all the links in the HTML
func TestGetLinksEndpoint(t *testing.T) {
	assert.Equal(t, "GetLinks", "GetLinksEndpoint", "TODO: Create get links test and endpoint")
}
