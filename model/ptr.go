package model

import "time"

type Ptr struct {
	Address   string    `json:"address"`
	Record    string    `json:"record"`
	Timestamp time.Time `json:"timestamp"`
}

type Ptrs []Ptr
