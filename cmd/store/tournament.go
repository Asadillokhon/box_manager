package store

import (
	"box-manager/cmd/model"
	"fmt"
)

func (s *MemoryStore) CreateTournament(t model.Tournament) (model.Tournament, error) {
	t.ID = s.nextId
	s.nextId++
	s.tournaments[t.ID] = t
	return t, nil
}
func (s *MemoryStore) GetTournament(id int) (model.Tournament, error) {
	tour, ok := s.tournaments[id]
	if !ok {
		return model.Tournament{}, fmt.Errorf("tournament with id %d not found", id)
	}
	return tour, nil
}
func (s *MemoryStore) UpdateTournament(id int, t model.Tournament) error {
	_, ok := s.tournaments[id]
	if !ok {
		return fmt.Errorf("tournament with id %d not found", id)
	}
	t.ID = id
	s.tournaments[id] = t
	return nil
}
func (s *MemoryStore) DeleteTournament(id int) error {
	_, ok := s.tournaments[id]
	if !ok {
		return fmt.Errorf("tournament with id %d not found", id)
	}
	delete(s.tournaments, id)
	return nil
}
func (s *MemoryStore) ListTournaments() []model.Tournament {
	result := make([]model.Tournament, 0, len(s.tournaments))
	for _, val := range s.tournaments {
		result = append(result, val)
	}
	return result
}
func (s *MemoryStore) AddParticipantToTournament(tournamentID, fighterID, place int) error {
	tour, ok := s.tournaments[tournamentID]
	if !ok {
		return fmt.Errorf("tournament with id %d not found", tournamentID)
	}
	result := make(map[int]struct{})
	for _, p := range tour.Participants {
		result[p.FighterID] = struct{}{}
	}

	if _, ok := result[fighterID]; ok {
		return fmt.Errorf("fighter with id %d already in tournament", fighterID)
	}
	newParticipant := model.TournamentParticipant{
		FighterID: fighterID,
		Place:     place,
	}
	tour.Participants = append(tour.Participants, newParticipant)
	s.tournaments[tournamentID] = tour
	return nil
}
func (s *MemoryStore) RemoveParticipantFromTournament(tournamentID, fighterID int) error {
	tour, ok := s.tournaments[tournamentID]
	if !ok {
		return fmt.Errorf("tournament with id %d not found", tournamentID)
	}
	idx := -1
	for i, p := range tour.Participants {
		if p.FighterID == fighterID {
			idx = i
			break
		}
	}
	if idx == -1 {
		return fmt.Errorf("participant with fighter id %d not found in tournament", fighterID)
	}
	tour.Participants = append(tour.Participants[:idx], tour.Participants[idx+1:]...)
	s.tournaments[tournamentID] = tour
	return nil
}
