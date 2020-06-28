package main

type PtrRepository interface {
	GetPtrs() (Ptrs, error)
}
