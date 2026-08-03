package store

import (
	"box-manager/cmd/model"
	"fmt"
)

func (s *MemoryStore) CreateFighter(f model.Fighter) (model.Fighter, error) {
	f.ID = s.nextId
	s.nextId++
	s.fighters[f.ID] = f
	return f, nil
}
func (s *MemoryStore) GetFighter(id int) (model.Fighter, error) {
	fighter, ok := s.fighters[id]
	if ok {
		return fighter, nil
	}
	return model.Fighter{}, fmt.Errorf("fighter with id %d not found", id)
}
func (s *MemoryStore) UpdateFighter(id int, f model.Fighter) error {
	_, ok := s.fighters[id]
	if !ok {
		return fmt.Errorf("fighter with id %d not found", id)
	}
	f.ID = id
	s.fighters[id] = f
	return nil
}
func (s *MemoryStore) DeleteFighter(id int) error {
	_, ok := s.fighters[id]
	if !ok {
		return fmt.Errorf("fighter with id %d not found", id)
	}
	delete(s.fighters, id)
	return nil
}
func (s *MemoryStore) ListFighters() []model.Fighter {
	result := make([]model.Fighter, 0, len(s.clubs))
	for _, val := range s.fighters {
		result = append(result, val)
	}
	return result

}
