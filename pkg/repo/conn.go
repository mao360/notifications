package repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func ConnToDB(config string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), config)
	if err != nil {
		return nil, err
	}
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		err = pool.Ping(context.Background())
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	return pool, err
}
