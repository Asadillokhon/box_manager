package store

import (
	"box-manager/cmd/model"
	"fmt"
)

func (s *MemoryStore) CreateFight(f model.Fight) (model.Fight, error) {
	_, err := s.GetFighter(f.Fighter1ID)
	if err != nil {
		return model.Fight{}, fmt.Errorf("fighter1: %w", err)

	}
	_, err = s.GetFighter(f.Fighter2ID)
	if err != nil {
		return model.Fight{}, fmt.Errorf("fighter2: %w", err)
	}
	if f.Fighter1ID == f.Fighter2ID {
		return model.Fight{}, fmt.Errorf("fighter cannot fight themselves")
	}

	f.ID = s.nextId
	s.nextId++
	s.fights[f.ID] = f
	return f, nil
}
func (s *MemoryStore) GetFight(id int) (model.Fight, error) {
	figh, ok := s.fights[id]
	if !ok {
		return model.Fight{}, fmt.Errorf("fight with id %d not found", id)
	}
	return figh, nil
}
func (s *MemoryStore) UpdateFight(id int, f model.Fight) error {
	_, ok := s.fights[id]
	if !ok {
		return fmt.Errorf("fight with id %d not found", id)
	}
	f.ID = id
	s.fights[id] = f
	return nil
}
func (s *MemoryStore) DeleteFight(id int) error {
	_, ok := s.fights[id]
	if !ok {
		return fmt.Errorf("fight with id %d not found", id)
	}
	delete(s.fights, id)
	return nil
}
func (s *MemoryStore) ListFights() []model.Fight {
	result := make([]model.Fight, 0, len(s.fights))
	for _, val := range s.fights {
		result = append(result, val)
	}
	return result
}
