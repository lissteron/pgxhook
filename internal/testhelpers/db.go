package testhelpers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectWithRandomDB(ctx context.Context, database string) (*pgxpool.Pool, string, error) {
	dbClient, err := pgxpool.New(ctx, "postgres://postgres:postgres@postgres:5432/"+database)
	if err != nil {
		return nil, "", fmt.Errorf("connect to database: %w", err)
	}

	dbName := getDBName(database)

	if err := createDB(ctx, dbClient, dbName); err != nil {
		return nil, "", err
	}

	dbClient.Close()

	dbClient, err = pgxpool.New(ctx, "postgres://postgres:postgres@postgres:5432/"+dbName)
	if err != nil {
		return nil, "", fmt.Errorf("connect to database: %w", err)
	}

	return dbClient, dbName, nil
}

func createDB(ctx context.Context, db *pgxpool.Pool, dbName string) error {
	if _, err := db.Exec(ctx, `CREATE DATABASE `+dbName); err != nil {
		return fmt.Errorf("create db %w", err)
	}

	return nil
}

func getDBName(name string) string {
	const base = 32

	return name + "_" + strconv.FormatInt(time.Now().UnixNano(), base)
}
