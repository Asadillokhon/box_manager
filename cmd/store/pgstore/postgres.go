package pgstore

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStore struct {
	db *pgxpool.Pool
}

// NewPostgresStore создает подключение к БД
func NewPostgresStore(ctx context.Context, dsn string) (*PostgresStore, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	// Проверяем связь
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}
	return &PostgresStore{db: pool}, nil
}

// Close обязательно вызывать defer при завершении программы
func (s *PostgresStore) Close() {
	s.db.Close()
}
