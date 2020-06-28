package repo

import "git.rappet.de/rappet/netmap-api/model"

type GetPtrsParameters struct {
	Like string
}

type PtrRepository interface {
	GetPtrs(parameters GetPtrsParameters) (model.Ptrs, error)
	GetPtr(address string) (*model.Ptr, error)
}
