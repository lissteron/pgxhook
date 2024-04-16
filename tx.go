package pgxhook

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type HookTx struct {
	tx   pgx.Tx
	conn *HookConn

	beforeHooks []BeforeHook
	afterHooks  []AfterHook
}

func newHookTx(ctx context.Context, conn *HookConn, txOptions *pgx.TxOptions) (*HookTx, error) {
	if txOptions == nil {
		txOptions = &pgx.TxOptions{}
	}

	tx, err := conn.conn.BeginTx(ctx, *txOptions)
	if err != nil {
		return nil, err
	}

	return &HookTx{
		tx:          tx,
		conn:        conn,
		beforeHooks: conn.beforeHooks,
		afterHooks:  conn.afterHooks,
	}, nil
}

func (t *HookTx) Begin(_ context.Context) (Tx, error) {
	return t, nil
}

func (t *HookTx) Commit(ctx context.Context) error {
	var err error

	for _, hook := range t.beforeHooks {
		if ctx, err = hook.Before(ctx, &HookData{
			Caller: CallerCommit,
		}); err != nil {
			return err
		}
	}

	err = t.tx.Commit(ctx)

	for _, hook := range t.afterHooks {
		if ctx, err = hook.After(ctx, &HookData{
			Caller: CallerCommit,
			Error:  err,
		}); err != nil {
			return err
		}
	}

	if err != nil {
		return err
	}

	return nil
}

func (t *HookTx) Rollback(ctx context.Context) error {
	var err error

	for _, hook := range t.beforeHooks {
		if ctx, err = hook.Before(ctx, &HookData{
			Caller: CallerRollback,
		}); err != nil {
			return err
		}
	}

	err = t.tx.Rollback(ctx)

	for _, hook := range t.afterHooks {
		if ctx, err = hook.After(ctx, &HookData{
			Caller: CallerRollback,
			Error:  err,
		}); err != nil {
			return err
		}
	}

	if err != nil {
		return err
	}

	return nil
}

func (t *HookTx) CopyFrom(
	ctx context.Context,
	tableName pgx.Identifier,
	columnNames []string,
	rowSrc pgx.CopyFromSource,
) (int64, error) {
	return t.tx.CopyFrom(ctx, tableName, columnNames, rowSrc)
}

func (t *HookTx) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	var err error

	for _, hook := range t.beforeHooks {
		if ctx, err = hook.Before(ctx, &HookData{
			Caller: CallerSendBatch,
			InTx:   true,
		}); err != nil {
			return &batchWithError{err: err}
		}
	}

	resp := t.tx.SendBatch(ctx, b)

	for _, hook := range t.afterHooks {
		if ctx, err = hook.After(ctx, &HookData{
			Caller: CallerSendBatch,
			Error:  err,
			InTx:   true,
		}); err != nil {
			return &batchWithError{err: err}
		}
	}

	return resp
}

func (t *HookTx) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	var err error

	for _, hook := range t.beforeHooks {
		if ctx, err = hook.Before(ctx, &HookData{
			Caller: CallerExec,
			Query:  sql,
			Args:   args,
			InTx:   true,
		}); err != nil {
			return pgconn.CommandTag{}, err
		}
	}

	resp, err := t.tx.Exec(ctx, sql, args...)

	for _, hook := range t.afterHooks {
		if ctx, err = hook.After(ctx, &HookData{
			Caller: CallerExec,
			Query:  sql,
			Args:   args,
			Error:  err,
			InTx:   true,
		}); err != nil {
			return pgconn.CommandTag{}, err
		}
	}

	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return resp, nil
}

func (t *HookTx) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	var err error

	for _, hook := range t.beforeHooks {
		if ctx, err = hook.Before(ctx, &HookData{
			Caller: CallerQuery,
			Query:  sql,
			Args:   args,
			InTx:   true,
		}); err != nil {
			return nil, err
		}
	}

	resp, err := t.tx.Query(ctx, sql, args...)

	for _, hook := range t.afterHooks {
		if ctx, err = hook.After(ctx, &HookData{
			Caller: CallerQuery,
			Query:  sql,
			Args:   args,
			Error:  err,
			InTx:   true,
		}); err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *HookTx) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	var err error

	for _, hook := range t.beforeHooks {
		if ctx, err = hook.Before(ctx, &HookData{
			Caller: CallerQueryRow,
			Query:  sql,
			Args:   args,
			InTx:   true,
		}); err != nil {
			return &scanWithError{err: err}
		}
	}

	resp := t.tx.QueryRow(ctx, sql, args...)

	for _, hook := range t.afterHooks {
		if ctx, err = hook.After(ctx, &HookData{
			Caller: CallerQueryRow,
			Query:  sql,
			Args:   args,
			Error:  err,
			InTx:   true,
		}); err != nil {
			return &scanWithError{err: err}
		}
	}

	return resp
}

func (t *HookTx) LargeObjects() pgx.LargeObjects {
	return t.tx.LargeObjects()
}

func (t *HookTx) Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error) {
	return t.tx.Prepare(ctx, name, sql)
}

// TODO remove
func (t *HookTx) Conn() *pgx.Conn {
	c, _ := t.conn.conn.Acquire(context.Background())

	return c.Conn()
}
