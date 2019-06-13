package pkg

import "sync"

type Inmem struct {
	mtx sync.Mutex
	store map[string]Book
}

func NewInmemStorage() *Inmem {
	return &Inmem{
		store: map[string]Book{},
	}
}
