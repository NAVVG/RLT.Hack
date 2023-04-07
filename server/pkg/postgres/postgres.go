package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Instance struct {
	Pool *pgxpool.Pool
}

func New(url string) (*Instance, error) {
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	db := &Instance{
		Pool: pool,
	}

	return db, nil
}
