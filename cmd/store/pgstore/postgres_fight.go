package pgstore

import (
	"box-manager/cmd/model"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (s *PostgresStore) UpdateFight(ctx context.Context, id int, f model.Fight) error {
	return nil
}
func (s *PostgresStore) DeleteFight(ctx context.Context, id int) error {
	return nil
}
func (s *PostgresStore) ListFights(ctx context.Context) []model.Fight {
	return []model.Fight{}
}

func (s *PostgresStore) CreateFight(ctx context.Context, f model.Fight) (model.Fight, error) {
	// 1. Создаем сам бой в таблице fights
	queryFight := `
		INSERT INTO fights (fighter1_id, fighter2_id, result) 
		VALUES ($1, $2, $3) 
		RETURNING id
	`
	var fightID int
	err := s.db.QueryRow(ctx, queryFight, f.Fighter1ID, f.Fighter2ID, f.Result).Scan(&fightID)
	if err != nil {
		return model.Fight{}, fmt.Errorf("failed to create fight: %w", err)
	}
	f.ID = fightID

	// 2. Сохраняем раунды в таблицу fight_rounds (если они есть)
	for i, round := range f.Rounds {
		queryRound := `
			INSERT INTO fight_rounds (fight_id, round_number, fighter1_score, fighter2_score) 
			VALUES ($1, $2, $3, $4)
		`
		_, err := s.db.Exec(ctx, queryRound, fightID, i+1, round.Fighter1Score, round.Fighter2Score)
		if err != nil {
			return model.Fight{}, fmt.Errorf("failed to save round %d: %w", i+1, err)
		}
	}

	return f, nil
}

func (s *PostgresStore) GetFight(ctx context.Context, id int) (model.Fight, error) {
	var f model.Fight
	// 1. Получаем бой
	queryFight := `
		SELECT id, fighter1_id, fighter2_id, result 
		FROM fights 
		WHERE id = $1
	`
	err := s.db.QueryRow(ctx, queryFight, id).Scan(
		&f.ID, &f.Fighter1ID, &f.Fighter2ID, &f.Result,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.Fight{}, fmt.Errorf("fight with id %d not found", id)
		}
		return model.Fight{}, fmt.Errorf("failed to get fight: %w", err)
	}

	// 2. Получаем раунды
	queryRounds := `
		SELECT round_number, fighter1_score, fighter2_score 
		FROM fight_rounds 
		WHERE fight_id = $1 
		ORDER BY round_number
	`
	rows, err := s.db.Query(ctx, queryRounds, id)
	if err != nil {
		return model.Fight{}, fmt.Errorf("failed to get rounds: %w", err)
	}
	defer rows.Close()

	var rounds []model.RoundsScore
	for rows.Next() {
		var r model.RoundsScore
		var roundNum int // нам не нужно сохранять номер раунда в модель, но мы его сканируем для порядка
		if err := rows.Scan(&roundNum, &r.Fighter1Score, &r.Fighter2Score); err == nil {
			rounds = append(rounds, r)
		}
	}
	f.Rounds = rounds

	return f, nil
}
