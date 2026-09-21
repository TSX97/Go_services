package db

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
)



func NewPool() (*pgxpool.Pool, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, "postgres://postgres:5432@database/postgtes")
	if err != nil {
		return nil, err
	}
	fmt.Println("connect to postgres")
	return pool, err
}
