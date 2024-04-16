package pgxhook

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type batchWithError struct {
	err error
}

func (b *batchWithError) Exec() (pgconn.CommandTag, error) {
	return pgconn.CommandTag{}, b.err
}

func (b *batchWithError) Query() (pgx.Rows, error) {
	return nil, b.err
}

func (b *batchWithError) QueryRow() pgx.Row {
	return &scanWithError{err: b.err}
}

func (b *batchWithError) Close() error {
	return b.err
}
