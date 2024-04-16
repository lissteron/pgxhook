//go:build integration

package pgxhook_test

import (
	"context"
	"testing"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/lissteron/pgxhook"
	"github.com/lissteron/pgxhook/gen/mocks"
	"github.com/lissteron/pgxhook/internal/testhelpers"
)

const (
	_defaultDBName = "postgres"

	_suiteTestTimeout = 10 * time.Second
)

const _testTable = `
	CREATE TABLE IF NOT EXISTS users (
		id                  SERIAL               PRIMARY KEY,
		name 			    VARCHAR(255)         NOT NULL,
		lastname 		    VARCHAR(255)         NOT NULL,
		email               VARCHAR(255)         NOT NULL,
		created_at          TIMESTAMPTZ          NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at          TIMESTAMPTZ          NOT NULL DEFAULT CURRENT_TIMESTAMP
	)
`

type userModel struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	LastName  string    `db:"lastname"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type pgxHookSuite struct {
	suite.Suite
	dbClient *pgxpool.Pool
	dbName   string
}

func TestPGXHookSuite(t *testing.T) {
	suite.Run(t, &pgxHookSuite{})
}

func (s *pgxHookSuite) SetupSuite() {
	var (
		err error
		ctx = context.Background()
	)

	s.dbClient, s.dbName, err = testhelpers.ConnectWithRandomDB(ctx, _defaultDBName)
	s.Require().NoError(err)

	_, err = s.dbClient.Exec(ctx, _testTable)
	s.Require().NoError(err)
}

func (s *pgxHookSuite) TearDownSuite() {
	s.dbClient.Close()
}

func (s *pgxHookSuite) BeforeTest(suiteName, testName string) {
	ctx, cancel := context.WithTimeout(context.Background(), _suiteTestTimeout)
	defer cancel()

	_, err := s.dbClient.Exec(ctx, `DELETE FROM users`)
	s.Require().NoError(err)
}

func (s *pgxHookSuite) TestBase() {
	ctx, cancel := context.WithTimeout(context.Background(), _suiteTestTimeout)
	defer cancel()

	conn := pgxhook.NewHookConn(s.dbClient)

	_, err := conn.Exec(ctx, `INSERT INTO users(name, lastname, email) VALUES($1, $2, $3)`, "John", "Doe", "test@example.com")
	s.Require().NoError(err)
}

func (s *pgxHookSuite) TestWithScany() {
	ctx, cancel := context.WithTimeout(context.Background(), _suiteTestTimeout)
	defer cancel()

	conn := pgxhook.NewHookConn(s.dbClient)

	user := userModel{
		Name:     "John",
		LastName: "Doe",
		Email:    "test@example.com",
	}

	_, err := conn.Exec(ctx, `INSERT INTO users(name, lastname, email) VALUES($1, $2, $3)`, user.Name, user.LastName, user.Email)
	s.Require().NoError(err)

	var got []userModel
	err = pgxscan.Select(ctx, conn, &got, `SELECT id, name, lastname, email, created_at, updated_at FROM users WHERE email = $1`, user.Email)
	s.Require().NoError(err)
	s.Require().NotZero(got[0].ID)
	s.Require().NotZero(got[0].CreatedAt)
	s.Require().NotZero(got[0].UpdatedAt)

	user.ID = got[0].ID
	user.CreatedAt = got[0].CreatedAt
	user.UpdatedAt = got[0].UpdatedAt

	s.Require().Equal([]userModel{user}, got)
}

func (s *pgxHookSuite) TestWithHook() {
	ctx, cancel := context.WithTimeout(context.Background(), _suiteTestTimeout)
	defer cancel()

	var (
		beforeHook = mocks.NewMockBeforeHook(s.T())
		afterHook  = mocks.NewMockAfterHook(s.T())
		conn       = pgxhook.NewHookConn(
			s.dbClient,
			pgxhook.WithBeforeHooks(beforeHook),
			pgxhook.WithAfterHooks(afterHook),
		)

		user = userModel{
			Name:     "John",
			LastName: "Doe",
			Email:    "test@example.com",
		}
	)

	beforeHook.EXPECT().Before(mock.Anything, &pgxhook.HookData{
		Query:  `INSERT INTO users(name, lastname, email) VALUES($1, $2, $3)`,
		Args:   []any{user.Name, user.LastName, user.Email},
		Caller: pgxhook.CallerExec,
	}).Once().Return(ctx, nil)

	afterHook.EXPECT().After(mock.Anything, &pgxhook.HookData{
		Query:  `INSERT INTO users(name, lastname, email) VALUES($1, $2, $3)`,
		Args:   []any{user.Name, user.LastName, user.Email},
		Caller: pgxhook.CallerExec,
		Error:  nil,
	}).Once().Return(ctx, nil)

	_, err := conn.Exec(ctx, `INSERT INTO users(name, lastname, email) VALUES($1, $2, $3)`, user.Name, user.LastName, user.Email)
	s.Require().NoError(err)

	beforeHook.EXPECT().Before(mock.Anything, &pgxhook.HookData{
		Query:  `SELECT id, name, lastname, email, created_at, updated_at FROM users WHERE email = $1`,
		Args:   []any{user.Email},
		Caller: pgxhook.CallerQuery,
	}).Once().Return(ctx, nil)

	afterHook.EXPECT().After(mock.Anything, &pgxhook.HookData{
		Query:  `SELECT id, name, lastname, email, created_at, updated_at FROM users WHERE email = $1`,
		Args:   []any{user.Email},
		Caller: pgxhook.CallerQuery,
		Error:  nil,
	}).Once().Return(ctx, nil)

	var got []userModel
	err = pgxscan.Select(ctx, conn, &got, `SELECT id, name, lastname, email, created_at, updated_at FROM users WHERE email = $1`, user.Email)
	s.Require().NoError(err)
	s.Require().NotZero(got[0].ID)
	s.Require().NotZero(got[0].CreatedAt)
	s.Require().NotZero(got[0].UpdatedAt)

	user.ID = got[0].ID
	user.CreatedAt = got[0].CreatedAt
	user.UpdatedAt = got[0].UpdatedAt

	s.Require().Equal([]userModel{user}, got)
}
