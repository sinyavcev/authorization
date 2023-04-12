package postgres

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	pgxpool  *pgxpool.Pool
	squirrel squirrel.StatementBuilderType
}

func NewPostgres(connString string) (*postgres, error) {
	sqlx := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	db, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}
	newPostgres := &postgres{db, sqlx}
	return newPostgres, nil
}
