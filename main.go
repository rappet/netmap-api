package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("starting netmap-api")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
