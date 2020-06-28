package main

import (
	"git.rappet.de/rappet/netmap-api/routing"
	"git.rappet.de/rappet/netmap-api/repo"
	"git.rappet.de/rappet/netmap-api/repo_mock"
	"log"
	"net/http"
)

var ptrRepository repo.PtrRepository = repo_mock.NewPtrRepository()

func main() {
	log.Println("starting netmap-api")

	router := routing.NewRouter(routes)

	log.Fatal(http.ListenAndServe(":8080", router))
}
