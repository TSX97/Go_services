package db

import (
	"fmt"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Ping(ctx context.Context, pool *pgxpool.Pool) error {
	fmt.Println("ping")
	return pool.Ping(ctx)
}
