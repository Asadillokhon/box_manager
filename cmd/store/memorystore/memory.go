package memorystore

import (
	"box-manager/cmd/model"
)

type MemoryStore struct {
	fighters    map[int]model.Fighter
	clubs       map[int]model.Club
	tournaments map[int]model.Tournament
	fights      map[int]model.Fight
	nextId      int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		fighters:    make(map[int]model.Fighter),
		clubs:       make(map[int]model.Club),
		tournaments: make(map[int]model.Tournament),
		fights:      make(map[int]model.Fight),
		nextId:      1,
	}
}
