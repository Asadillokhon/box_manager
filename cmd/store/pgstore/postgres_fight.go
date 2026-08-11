package pgstore

import (
	"box-manager/cmd/model"
	"context"
)

func (s *PostgresStore) CreateFight(ctx context.Context, f model.Fight) (model.Fight, error) {
	return f, nil
}
func (s *PostgresStore) GetFight(ctx context.Context, id int) (model.Fight, error) {
	return model.Fight{}, nil
}
func (s *PostgresStore) UpdateFight(ctx context.Context, id int, f model.Fight) error {
	return nil
}
func (s *PostgresStore) DeleteFight(ctx context.Context, id int) error {
	return nil
}
func (s *PostgresStore) ListFights(ctx context.Context) []model.Fight {
	return []model.Fight{}
}
