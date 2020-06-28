package repo_mock

import (
	"git.rappet.de/rappet/netmap-api/model"
	"time"
)

type PtrRepository struct {
	ptrs map[string]model.Ptr
}

func NewPtrRepository() PtrRepository {
	ptrs := make(map[string]model.Ptr)
	ptrs["10.0.0.1"] = model.Ptr{Address: "10.0.0.1", Record: "host1.example.com", Timestamp: time.Now()}
	ptrs["10.0.0.2"] = model.Ptr{Address: "10.0.0.1", Record: "host1.example.com", Timestamp: time.Now()}
	ptrs["10.0.0.3"] = model.Ptr{Address: "10.0.0.1", Record: "host1.example.com", Timestamp: time.Now()}
	return PtrRepository{
		ptrs: ptrs,
	}
}

func (mock PtrRepository) GetPtrs() (model.Ptrs, error) {
	ptrs := make([]model.Ptr, 0)
	for _, ptr := range mock.ptrs {
		ptrs = append(ptrs, ptr)
	}

	return ptrs, nil
}