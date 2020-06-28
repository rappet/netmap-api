package main

import "time"

type PtrRepositoryMock struct {
	ptrs map[string]Ptr
}

func NewPtrRepositoryMock() PtrRepositoryMock {
	ptrs := make(map[string]Ptr)
	ptrs["10.0.0.1"] = Ptr{Address: "10.0.0.1", Record: "host1.example.com", Timestamp: time.Now()}
	ptrs["10.0.0.2"] = Ptr{Address: "10.0.0.1", Record: "host1.example.com", Timestamp: time.Now()}
	ptrs["10.0.0.3"] = Ptr{Address: "10.0.0.1", Record: "host1.example.com", Timestamp: time.Now()}
	return PtrRepositoryMock{
		ptrs: ptrs,
	}
}

func (mock PtrRepositoryMock) GetPtrs() (Ptrs, error) {
	ptrs := make([]Ptr, 0)
	for _, ptr := range mock.ptrs {
		ptrs = append(ptrs, ptr)
	}

	return ptrs, nil
}