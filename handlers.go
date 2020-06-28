package main

import (
	"git.rappet.de/rappet/netmap-api/routing"
	"github.com/gorilla/mux"
	"net/http"
)

func PtrIndex(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ptrs, err := ptrRepository.GetPtrs()
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