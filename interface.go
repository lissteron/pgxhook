package pgxhook

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Tx = pgx.Tx

type Conn interface {
	Begin(ctx context.Context) (Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (Tx, error)

	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	SendBatch(ctx context.Context, b *pgx.Batch) (br pgx.BatchResults)

	Ping(ctx context.Context) error
	Close()

	Config() *pgxpool.Config
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
}

type BeforeHook interface {
	Before(ctx context.Context, input *HookData) (context.Context, error)
}

type AfterHook interface {
	After(ctx context.Context, input *HookData) (context.Context, error)
}

type FullHook interface {
	BeforeHook
	AfterHook
}
