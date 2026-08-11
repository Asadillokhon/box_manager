package memorystore

import (
	"box-manager/cmd/model"
	"context"
	"fmt"
)

func (s *MemoryStore) CreateClub(ctx context.Context, c model.Club) (model.Club, error) {
	c.ID = s.nextId
	s.nextId++
	s.clubs[c.ID] = c
	return c, nil
}
func (s *MemoryStore) GetClub(ctx context.Context, id int) (model.Club, error) {
	club, ok := s.clubs[id]
	if ok {
		return club, nil
	}
	return model.Club{}, fmt.Errorf("Club with id %d not found", id)
}
func (s *MemoryStore) DeleteClub(ctx context.Context, id int) error {
	_, ok := s.clubs[id]
	if !ok {
		return fmt.Errorf("Club with id %d not found", id)
	}
	delete(s.clubs, id)
	return nil
}
func (s *MemoryStore) UpdateClub(ctx context.Context, id int, c model.Club) error {
	_, ok := s.clubs[id]
	if !ok {
		return fmt.Errorf("Club with id %d not found", id)
	}
	c.ID = id
	s.clubs[id] = c
	return nil
}
func (s *MemoryStore) ListClubs(ctx context.Context) []model.Club {
	result := make([]model.Club, 0, len(s.clubs))
	for _, val := range s.clubs {
		result = append(result, val)
	}
	return result
}
