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

const (
	userEmail string = "demo@usuario.com"
	password  string = "pipjY7-guknaq-nancex"
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
	values := map[string]string{
		"userEmail": userEmail,
		"password":  password,
	}
	json_data, err := json.Marshal(values)
	handleErr(err)

	resp, err := http.NewRequest("POST", "/login", bytes.NewBuffer(json_data))
	handleErr(err)

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	assert.Equal(t, "", res["json"], "TODO: Fix test and create endpoint")
}
