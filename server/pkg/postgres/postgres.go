package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func New(url string) (*Postgres, error) {
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	db := &Postgres{
		Pool: pool,
	}

	return db, nil
}
