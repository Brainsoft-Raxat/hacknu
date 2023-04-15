package repository

import (
	"context"
	"disease-api/internal/app/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

type postgres struct {
	db  *pgxpool.Pool
	cfg config.Postgres
}

// всё взаимодействие с постгрес
func (r *postgres) DoSomething(ctx context.Context, name string) (value string, err error) {
	//err = r.db.QueryRow(ctx, q, c.Cname, c.Population).Scan(&cname)
	//if err != nil {
	//	return
	//}

	return "ok", nil
}

func NewPostgresRepository(db *pgxpool.Pool, cfg config.Postgres) Postgres {
	return &postgres{
		db:  db,
		cfg: cfg,
	}
}
