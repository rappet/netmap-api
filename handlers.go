package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func PtrIndex(w http.ResponseWriter, r *http.Request) {
	ptrs, err := ptrRepository.GetPtrs()
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Server error")
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ptrs); err != nil {
		log.Fatalln(err)
	}
}