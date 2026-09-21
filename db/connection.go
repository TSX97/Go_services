package db

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
	"os"
)



func NewPool() (*pgxpool.Pool, error) {
	ctx := context.Background()
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	config := fmt.Sprintf("postgres://%s:%s@%s/%s", user, pass, host, port, name)

	pool, err := pgxpool.New(ctx, config)
	if err != nil {
		return nil, err
	}
	fmt.Println("connect to postgres")
	return pool, err
}
