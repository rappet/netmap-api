package main

import (
	"git.rappet.de/rappet/netmap-api/repo"
	"git.rappet.de/rappet/netmap-api/routing"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func PtrIndex(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var likeParam string
	likeParams, ok := r.URL.Query()["like"]
	if ok && len(likeParams) > 0 {
		likeParam = likeParams[0]
	}

	parameters := repo.GetPtrsParameters{Like: likeParam}
	log.Print(parameters)

	ptrs, err := ptrRepository.GetPtrs(parameters)
	if err != nil {
		return nil, err
	}

	return ptrs, nil
}

func GetPtr(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	address := vars["address"]
	ptr, err := ptrRepository.GetPtr(address)
	if err != nil {
		return nil, err
	}
	if ptr == nil {
		return nil, routing.HttpError{
			Message: "PTR not found",
			Code:    http.StatusNotFound,
		}
	}
	return ptr, err
}