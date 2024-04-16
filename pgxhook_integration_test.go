//go:build integration

package pgxhook_test

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"

	"github.com/lissteron/pgxhook"
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
