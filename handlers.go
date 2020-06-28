package main

import (
	"net/http"
)

func PtrIndex(r *http.Request) (interface{}, error) {
	ptrs, err := ptrRepository.GetPtrs()
	if err != nil {
		return nil, err
	}

	return ptrs, nil
}