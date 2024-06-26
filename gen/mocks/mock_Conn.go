// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	pgconn "github.com/jackc/pgx/v5/pgconn"
	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v5"

	pgxhook "github.com/lissteron/pgxhook"
)

// MockConn is an autogenerated mock type for the Conn type
type MockConn struct {
	mock.Mock
}

type MockConn_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConn) EXPECT() *MockConn_Expecter {
	return &MockConn_Expecter{mock: &_m.Mock}
}

// Begin provides a mock function with given fields: ctx
func (_m *MockConn) Begin(ctx context.Context) (pgx.Tx, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Begin")
	}

	var r0 pgx.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (pgx.Tx, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) pgx.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConn_Begin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Begin'
type MockConn_Begin_Call struct {
	*mock.Call
}

// Begin is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockConn_Expecter) Begin(ctx interface{}) *MockConn_Begin_Call {
	return &MockConn_Begin_Call{Call: _e.mock.On("Begin", ctx)}
}

func (_c *MockConn_Begin_Call) Run(run func(ctx context.Context)) *MockConn_Begin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockConn_Begin_Call) Return(_a0 pgx.Tx, _a1 error) *MockConn_Begin_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConn_Begin_Call) RunAndReturn(run func(context.Context) (pgx.Tx, error)) *MockConn_Begin_Call {
	_c.Call.Return(run)
	return _c
}

// BeginTx provides a mock function with given fields: ctx, txOptions
func (_m *MockConn) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	ret := _m.Called(ctx, txOptions)

	if len(ret) == 0 {
		panic("no return value specified for BeginTx")
	}

	var r0 pgx.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.TxOptions) (pgx.Tx, error)); ok {
		return rf(ctx, txOptions)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.TxOptions) pgx.Tx); ok {
		r0 = rf(ctx, txOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.TxOptions) error); ok {
		r1 = rf(ctx, txOptions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConn_BeginTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BeginTx'
type MockConn_BeginTx_Call struct {
	*mock.Call
}

// BeginTx is a helper method to define mock.On call
//   - ctx context.Context
//   - txOptions pgx.TxOptions
func (_e *MockConn_Expecter) BeginTx(ctx interface{}, txOptions interface{}) *MockConn_BeginTx_Call {
	return &MockConn_BeginTx_Call{Call: _e.mock.On("BeginTx", ctx, txOptions)}
}

func (_c *MockConn_BeginTx_Call) Run(run func(ctx context.Context, txOptions pgx.TxOptions)) *MockConn_BeginTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(pgx.TxOptions))
	})
	return _c
}

func (_c *MockConn_BeginTx_Call) Return(_a0 pgx.Tx, _a1 error) *MockConn_BeginTx_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConn_BeginTx_Call) RunAndReturn(run func(context.Context, pgx.TxOptions) (pgx.Tx, error)) *MockConn_BeginTx_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields: ctx
func (_m *MockConn) Close(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockConn_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockConn_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockConn_Expecter) Close(ctx interface{}) *MockConn_Close_Call {
	return &MockConn_Close_Call{Call: _e.mock.On("Close", ctx)}
}

func (_c *MockConn_Close_Call) Run(run func(ctx context.Context)) *MockConn_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockConn_Close_Call) Return(_a0 error) *MockConn_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConn_Close_Call) RunAndReturn(run func(context.Context) error) *MockConn_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Conn provides a mock function with given fields:
func (_m *MockConn) Conn() pgxhook.InputConn {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Conn")
	}

	var r0 pgxhook.InputConn
	if rf, ok := ret.Get(0).(func() pgxhook.InputConn); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgxhook.InputConn)
		}
	}

	return r0
}

// MockConn_Conn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Conn'
type MockConn_Conn_Call struct {
	*mock.Call
}

// Conn is a helper method to define mock.On call
func (_e *MockConn_Expecter) Conn() *MockConn_Conn_Call {
	return &MockConn_Conn_Call{Call: _e.mock.On("Conn")}
}

func (_c *MockConn_Conn_Call) Run(run func()) *MockConn_Conn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConn_Conn_Call) Return(_a0 pgxhook.InputConn) *MockConn_Conn_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConn_Conn_Call) RunAndReturn(run func() pgxhook.InputConn) *MockConn_Conn_Call {
	_c.Call.Return(run)
	return _c
}

// Exec provides a mock function with given fields: ctx, sql, args
func (_m *MockConn) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, sql)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 pgconn.CommandTag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (pgconn.CommandTag, error)); ok {
		return rf(ctx, sql, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) pgconn.CommandTag); ok {
		r0 = rf(ctx, sql, args...)
	} else {
		r0 = ret.Get(0).(pgconn.CommandTag)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, sql, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConn_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockConn_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - sql string
//   - args ...interface{}
func (_e *MockConn_Expecter) Exec(ctx interface{}, sql interface{}, args ...interface{}) *MockConn_Exec_Call {
	return &MockConn_Exec_Call{Call: _e.mock.On("Exec",
		append([]interface{}{ctx, sql}, args...)...)}
}

func (_c *MockConn_Exec_Call) Run(run func(ctx context.Context, sql string, args ...interface{})) *MockConn_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockConn_Exec_Call) Return(_a0 pgconn.CommandTag, _a1 error) *MockConn_Exec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConn_Exec_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (pgconn.CommandTag, error)) *MockConn_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// Ping provides a mock function with given fields: ctx
func (_m *MockConn) Ping(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Ping")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockConn_Ping_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ping'
type MockConn_Ping_Call struct {
	*mock.Call
}

// Ping is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockConn_Expecter) Ping(ctx interface{}) *MockConn_Ping_Call {
	return &MockConn_Ping_Call{Call: _e.mock.On("Ping", ctx)}
}

func (_c *MockConn_Ping_Call) Run(run func(ctx context.Context)) *MockConn_Ping_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockConn_Ping_Call) Return(_a0 error) *MockConn_Ping_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConn_Ping_Call) RunAndReturn(run func(context.Context) error) *MockConn_Ping_Call {
	_c.Call.Return(run)
	return _c
}

// Query provides a mock function with given fields: ctx, sql, args
func (_m *MockConn) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, sql)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 pgx.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (pgx.Rows, error)); ok {
		return rf(ctx, sql, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) pgx.Rows); ok {
		r0 = rf(ctx, sql, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, sql, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConn_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type MockConn_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//   - ctx context.Context
//   - sql string
//   - args ...interface{}
func (_e *MockConn_Expecter) Query(ctx interface{}, sql interface{}, args ...interface{}) *MockConn_Query_Call {
	return &MockConn_Query_Call{Call: _e.mock.On("Query",
		append([]interface{}{ctx, sql}, args...)...)}
}

func (_c *MockConn_Query_Call) Run(run func(ctx context.Context, sql string, args ...interface{})) *MockConn_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockConn_Query_Call) Return(_a0 pgx.Rows, _a1 error) *MockConn_Query_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConn_Query_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (pgx.Rows, error)) *MockConn_Query_Call {
	_c.Call.Return(run)
	return _c
}

// QueryRow provides a mock function with given fields: ctx, sql, args
func (_m *MockConn) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	var _ca []interface{}
	_ca = append(_ca, ctx, sql)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryRow")
	}

	var r0 pgx.Row
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) pgx.Row); ok {
		r0 = rf(ctx, sql, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Row)
		}
	}

	return r0
}

// MockConn_QueryRow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryRow'
type MockConn_QueryRow_Call struct {
	*mock.Call
}

// QueryRow is a helper method to define mock.On call
//   - ctx context.Context
//   - sql string
//   - args ...interface{}
func (_e *MockConn_Expecter) QueryRow(ctx interface{}, sql interface{}, args ...interface{}) *MockConn_QueryRow_Call {
	return &MockConn_QueryRow_Call{Call: _e.mock.On("QueryRow",
		append([]interface{}{ctx, sql}, args...)...)}
}

func (_c *MockConn_QueryRow_Call) Run(run func(ctx context.Context, sql string, args ...interface{})) *MockConn_QueryRow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockConn_QueryRow_Call) Return(_a0 pgx.Row) *MockConn_QueryRow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConn_QueryRow_Call) RunAndReturn(run func(context.Context, string, ...interface{}) pgx.Row) *MockConn_QueryRow_Call {
	_c.Call.Return(run)
	return _c
}

// SendBatch provides a mock function with given fields: ctx, b
func (_m *MockConn) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	ret := _m.Called(ctx, b)

	if len(ret) == 0 {
		panic("no return value specified for SendBatch")
	}

	var r0 pgx.BatchResults
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Batch) pgx.BatchResults); ok {
		r0 = rf(ctx, b)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.BatchResults)
		}
	}

	return r0
}

// MockConn_SendBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendBatch'
type MockConn_SendBatch_Call struct {
	*mock.Call
}

// SendBatch is a helper method to define mock.On call
//   - ctx context.Context
//   - b *pgx.Batch
func (_e *MockConn_Expecter) SendBatch(ctx interface{}, b interface{}) *MockConn_SendBatch_Call {
	return &MockConn_SendBatch_Call{Call: _e.mock.On("SendBatch", ctx, b)}
}

func (_c *MockConn_SendBatch_Call) Run(run func(ctx context.Context, b *pgx.Batch)) *MockConn_SendBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*pgx.Batch))
	})
	return _c
}

func (_c *MockConn_SendBatch_Call) Return(br pgx.BatchResults) *MockConn_SendBatch_Call {
	_c.Call.Return(br)
	return _c
}

func (_c *MockConn_SendBatch_Call) RunAndReturn(run func(context.Context, *pgx.Batch) pgx.BatchResults) *MockConn_SendBatch_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockConn creates a new instance of MockConn. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockConn(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockConn {
	mock := &MockConn{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
