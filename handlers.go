package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func PtrIndex(w http.ResponseWriter, r *http.Request) {
	ptrs := Ptrs{
		Ptr{Address: "10.0.0.1", Record: "host1.example.net", Timestamp: time.Now()},
		Ptr{Address: "10.0.0.2", Record: "host2.example.net", Timestamp: time.Now()},
		Ptr{Address: "10.0.0.3", Record: "host3.example.net", Timestamp: time.Now()},
	}

	json.NewEncoder(w).Encode(ptrs)
}