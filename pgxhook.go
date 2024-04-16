package pgxhook

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HookConn struct {
	conn        *pgxpool.Pool
	beforeHooks []BeforeHook
	afterHooks  []AfterHook
}

func NewHookConn(conn *pgxpool.Pool, opts ...HookConnOption) *HookConn {
	c := &HookConn{conn: conn}

	for _, opt := range opts {
		opt.apply(c)
	}

	return c
}

func (c *HookConn) Begin(ctx context.Context) (Tx, error) {
	var err error

	for _, hook := range c.beforeHooks {
		if ctx, err = hook.Before(ctx, &HookData{
			Caller: CallerBegin,
		}); err != nil {
			return nil, err
		}
	}

	resp, err := newHookTx(ctx, c, nil)

	for _, hook := range c.afterHooks {
		if ctx, err = hook.After(ctx, &HookData{
			Caller: CallerBegin,
			Error:  err,
		}); err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *HookConn) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (Tx, error) {
	var err error

	for _, hook := range c.beforeHooks {
		if ctx, err = hook.Before(ctx, &HookData{
			Caller: CallerBegin,
		}); err != nil {
			return nil, err
		}
	}

	resp, err := newHookTx(ctx, c, &txOptions)

	for _, hook := range c.afterHooks {
		if ctx, err = hook.After(ctx, &HookData{
			Caller: CallerBegin,
			Error:  err,
		}); err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *HookConn) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	var err error

	for _, hook := range c.beforeHooks {
		if ctx, err = hook.Before(ctx, &HookData{
			Caller: CallerExec,
			Query:  sql,
			Args:   args,
		}); err != nil {
			return pgconn.CommandTag{}, err
		}
	}

	resp, err := c.conn.Exec(ctx, sql, args...)

	for _, hook := range c.afterHooks {
		if ctx, err = hook.After(ctx, &HookData{
			Caller: CallerExec,
			Query:  sql,
			Args:   args,
			Error:  err,
		}); err != nil {
			return pgconn.CommandTag{}, err
		}
	}

	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return resp, nil
}

func (c *HookConn) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	var err error

	for _, hook := range c.beforeHooks {
		if ctx, err = hook.Before(ctx, &HookData{
			Caller: CallerQuery,
			Query:  sql,
			Args:   args,
		}); err != nil {
			return nil, err
		}
	}

	resp, err := c.conn.Query(ctx, sql, args...)

	for _, hook := range c.afterHooks {
		if ctx, err = hook.After(ctx, &HookData{
			Caller: CallerQuery,
			Query:  sql,
			Args:   args,
			Error:  err,
		}); err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *HookConn) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	var err error

	for _, hook := range c.beforeHooks {
		if ctx, err = hook.Before(ctx, &HookData{
			Caller: CallerQueryRow,
			Query:  sql,
			Args:   args,
		}); err != nil {
			return &scanWithError{err: err}
		}
	}

	resp := c.conn.QueryRow(ctx, sql, args...)

	for _, hook := range c.afterHooks {
		if ctx, err = hook.After(ctx, &HookData{
			Caller: CallerQueryRow,
			Query:  sql,
			Args:   args,
			Error:  err,
		}); err != nil {
			return &scanWithError{err: err}
		}
	}

	return resp
}

func (c *HookConn) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	var err error

	for _, hook := range c.beforeHooks {
		if ctx, err = hook.Before(ctx, &HookData{
			Caller: CallerSendBatch,
		}); err != nil {
			return &batchWithError{err: err}
		}
	}

	resp := c.conn.SendBatch(ctx, b)

	for _, hook := range c.afterHooks {
		if ctx, err = hook.After(ctx, &HookData{
			Caller: CallerSendBatch,
			Error:  err,
		}); err != nil {
			return &batchWithError{err: err}
		}
	}

	return resp
}

func (c *HookConn) Ping(ctx context.Context) error {
	var err error

	for _, hook := range c.beforeHooks {
		if ctx, err = hook.Before(ctx, &HookData{
			Caller: CallerPing,
		}); err != nil {
			return err
		}
	}

	err = c.conn.Ping(ctx)

	for _, hook := range c.afterHooks {
		if ctx, err = hook.After(ctx, &HookData{
			Caller: CallerPing,
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

func (c *HookConn) Close() {
	c.conn.Close()
}

func (c *HookConn) Config() *pgxpool.Config {
	return c.conn.Config()
}

func (c *HookConn) CopyFrom(
	ctx context.Context,
	tableName pgx.Identifier,
	columnNames []string,
	rowSrc pgx.CopyFromSource,
) (int64, error) {
	return c.conn.CopyFrom(ctx, tableName, columnNames, rowSrc)
}
