package pgxhook

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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
	Close(ctx context.Context) error

	Conn() InputConn
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

type InputConn interface {
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	SendBatch(ctx context.Context, b *pgx.Batch) (br pgx.BatchResults)

	Begin(ctx context.Context) (Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (Tx, error)
	Ping(ctx context.Context) error
}

type ConnCloser interface {
	Close(ctx context.Context) error
}

type PoolCloser interface {
	Close()
}
