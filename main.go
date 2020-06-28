package main

import (
	"log"
	"net/http"
)

var ptrRepository PtrRepository = NewPtrRepositoryMock()

func main() {
	log.Println("starting netmap-api")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
