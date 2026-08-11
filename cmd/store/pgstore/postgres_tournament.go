package pgstore

import (
	"box-manager/cmd/model"
	"context"
)

func (s *PostgresStore) CreateTournament(ctx context.Context, t model.Tournament) (model.Tournament, error) {
	return t, nil
}
func (s *PostgresStore) GetTournament(ctx context.Context, id int) (model.Tournament, error) {
	return model.Tournament{}, nil
}

func (s *PostgresStore) UpdateTournament(ctx context.Context, id int, t model.Tournament) error {
	return nil
}

func (s *PostgresStore) DeleteTournament(ctx context.Context, id int) error {
	return nil
}
func (s *PostgresStore) ListTournaments(ctx context.Context) []model.Tournament {
	return []model.Tournament{}
}
func (s *PostgresStore) AddParticipantToTournament(ctx context.Context, tournamentID, fighterID, place int) error {
	return nil
}
func (s *PostgresStore) RemoveParticipantFromTournament(ctx context.Context, tournamentID, fighterID int) error {
	return nil
}
