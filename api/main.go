package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	return router
}

func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
	response.Write([]byte("Hello World"))
}

func main() {
	router := Router()
	log.Fatal(http.ListenAndServe(":12345", router))
}
