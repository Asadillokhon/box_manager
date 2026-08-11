package pgstore

import (
	"box-manager/cmd/model"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (s *PostgresStore) CreateClub(ctx context.Context, c model.Club) (model.Club, error) {
	query := `
	INSERT INTO clubs (name, address) 
	VALUES ($1, $2) 
	RETURNING id
	`
	var id int
	err := s.db.QueryRow(ctx, query, c.Name, c.Address).Scan(&id)
	if err != nil {
		return model.Club{}, fmt.Errorf("failed to create club: %w", err)
	}
	c.ID = id
	return c, nil
}
func (s *PostgresStore) GetClub(ctx context.Context, id int) (model.Club, error) {
	var c model.Club
	query := `
	SELECT id, name, address 
	FROM clubs 
	WHERE id = $1
	`
	err := s.db.QueryRow(ctx, query, id).Scan(&c.ID, &c.Name, &c.Address)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.Club{}, fmt.Errorf("Club with id %d not found", id)
		}
		return model.Club{}, fmt.Errorf("failed to get club: %w", err)
	}
	return c, nil
}

func (s *PostgresStore) UpdateClub(ctx context.Context, id int, c model.Club) error {
	query := `
	UPDATE clubs 
	SET name=$1, address=$2 
	WHERE id=$3
	`
	tag, err := s.db.Exec(ctx, query, c.Name, c.Address, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("Club with id %d not found", id)
	}
	return nil
}

func (s *PostgresStore) DeleteClub(ctx context.Context, id int) error {
	query := `
	DELETE FROM clubs 
	WHERE id = $1
	`
	tag, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("Club with id %d not found", id)
	}
	return nil
}

func (s *PostgresStore) ListClubs(ctx context.Context) []model.Club {
	query := `
	SELECT id, name, address 
	FROM clubs 
	ORDER BY id
	`
	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return []model.Club{}
	}
	defer rows.Close()

	var clubs []model.Club
	for rows.Next() {
		var c model.Club
		if err := rows.Scan(&c.ID, &c.Name, &c.Address); err == nil {
			clubs = append(clubs, c)
		}
	}
	return clubs
}
