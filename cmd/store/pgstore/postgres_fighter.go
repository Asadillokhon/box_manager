package pgstore

import (
	"box-manager/cmd/model"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (s *PostgresStore) CreateFighter(ctx context.Context, f model.Fighter) (model.Fighter, error) {
	query := `
	INSERT INTO fighters(first_name, last_name, birth_date, weight, category, club_id)
	VALUES($1, $2, $3, $4, $5, $6)
	RETURNING id
	`
	var id int
	err := s.db.QueryRow(ctx,
		query,
		f.FirstName,
		f.LastName,
		f.BirthDate,
		f.Weight,
		f.Category,
		f.ClubID,
	).Scan(&id)
	if err != nil {
		return model.Fighter{}, fmt.Errorf("failed to create fighter: %w", err)
	}
	f.ID = id
	return f, nil
}

func (s *PostgresStore) GetFighter(ctx context.Context, id int) (model.Fighter, error) {
	var f model.Fighter
	query := `
		SELECT id, first_name, last_name, birth_date, weight, category, club_id
		FROM fighters
		WHERE id = $1
	`
	err := s.db.QueryRow(ctx, query, id).Scan(
		&f.ID,
		&f.FirstName,
		&f.LastName,
		&f.BirthDate,
		&f.Weight,
		&f.Category,
		&f.ClubID,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.Fighter{}, fmt.Errorf("fighter with id %d not found", id)
		}
		return model.Fighter{}, fmt.Errorf("failed go get club: %w", err)
	}
	return f, nil
}

func (s *PostgresStore) UpdateFighter(ctx context.Context, id int, f model.Fighter) error {
	query := `
	UPDATE fighters
	SET firts_name= $1, last_name =$2, birth_date=$3, weight=$4, category=$5, club_id=$6
	WHERE id = $7
	`
	tag, err := s.db.Exec(ctx, query, f.FirstName, f.LastName, f.BirthDate, f.Weight, f.Category, f.ClubID, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("Fighter with id %d not found", id)
	}
	return nil
}

func (s *PostgresStore) DeleteFighter(ctx context.Context, id int) error {
	query := `
	DELETE FROM fighters
	WHERE id = $1
	`
	tag, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("fighter with id %d not found", id)
	}
	return nil
}

func (s *PostgresStore) ListFighters(ctx context.Context) []model.Fighter {
	query := `
	SELECT id, firts_name, last_name, birth_date, weight, category, club_id
	FROM fighters
	`
	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return []model.Fighter{}
	}
	defer rows.Close()
	var fighters []model.Fighter
	for rows.Next() {
		var f model.Fighter
		if err := rows.Scan(&f.ID,
			&f.FirstName,
			&f.LastName,
			&f.BirthDate,
			&f.Weight,
			&f.Category,
			&f.ClubID,
		); err == nil {
			fighters = append(fighters, f)
		}
	}
	return fighters
}
