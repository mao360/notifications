package repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnToDB(config string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), config)
	if err != nil {
		return nil, err
	}
	return pool, err
}
