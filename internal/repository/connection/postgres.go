package connection

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"hacknu/internal/app/config"
	"time"
)

func DialPostgres(ctx context.Context, cfg config.Postgres) (*pgxpool.Pool, error) {
	//dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%v/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	pool, err := pgxpool.Connect(ctx, cfg.URL)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
