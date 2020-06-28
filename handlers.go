package main

import (
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
	return ptr, err
}