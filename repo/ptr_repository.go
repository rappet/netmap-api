package repo

import "git.rappet.de/rappet/netmap-api/model"

type PtrRepository interface {
	GetPtrs() (model.Ptrs, error)
	GetPtr(address string) (*model.Ptr, error)
}
