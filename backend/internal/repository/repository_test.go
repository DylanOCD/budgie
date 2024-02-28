/*
 * Copyright (c) 2024 Dylan O' Connor Desmond
 */

package repository

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DylanOCD/budgie/backend/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type RepositorySuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo IRepository
}

func (s *RepositorySuite) SetupTest() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	dialector := postgres.New(
		postgres.Config{
			DSN:                  "sqlmock_db",
			DriverName:           "postgres",
			Conn:                 db,
			PreferSimpleProtocol: true,
		},
	)
	s.db, err = gorm.Open(dialector, &gorm.Config{})
	require.NoError(s.T(), err)

	r := New(s.db)
	s.repo = &r
}

func (s *RepositorySuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}

func (s *RepositorySuite) TestGetIncomes() {
	query := `SELECT * FROM "incomes" WHERE "incomes"."deleted_at" IS NULL`

	i1 := &models.Income{
		Model:   gorm.Model{ID: uint(1)},
		Account: "Bank of Ireland",
		Amount:  1234.50,
		Date:    time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC),
		Source:  "Work",
	}
	i2 := &models.Income{
		Model:   gorm.Model{ID: uint(2)},
		Account: "Bank of Ireland",
		Amount:  1234.50,
		Date:    time.Date(2024, 1, 15, 9, 0, 0, 0, time.UTC),
		Source:  "Work",
	}

	rows := s.mock.NewRows([]string{"id", "account", "amount", "date", "source"}).
		AddRow(i1.ID, i1.Account, i1.Amount, i1.Date, i1.Source).
		AddRow(i2.ID, i2.Account, i2.Amount, i2.Date, i2.Source)
	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	incomes, err := s.repo.GetIncomes()
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), incomes)
	assert.Equal(s.T(), i1.ID, incomes[0].ID, "Income 1 ID matches expected")
	assert.Equal(s.T(), i1.Account, incomes[0].Account, "Income 1 Account matches expected")
	assert.Equal(s.T(), i1.Amount, incomes[0].Amount, "Income 1 Amount matches expected")
	assert.Equal(s.T(), i1.Date, incomes[0].Date, "Income 1 Date matches expected")
	assert.Equal(s.T(), i1.Source, incomes[0].Source, "Income 1 Source matches expected")
	assert.Equal(s.T(), i2.ID, incomes[1].ID, "Income 2 ID matches expected")
	assert.Equal(s.T(), i2.Account, incomes[1].Account, "Income 2 Account matches expected")
	assert.Equal(s.T(), i2.Amount, incomes[1].Amount, "Income 2 Amount matches expected")
	assert.Equal(s.T(), i2.Date, incomes[1].Date, "Income 2 Date matches expected")
	assert.Equal(s.T(), i2.Source, incomes[1].Source, "Income 2 Source matches expected")
}

func (s *RepositorySuite) TestGetExpenses() {
	query := `SELECT * FROM "expenses" WHERE "expenses"."deleted_at" IS NULL`

	e1 := &models.Expense{
		Model:   gorm.Model{ID: uint(1)},
		Account: "Bank of Ireland",
		Amount:  1234.50,
		Date:    time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC),
		Vendor:  "Work",
	}
	e2 := &models.Expense{
		Model:   gorm.Model{ID: uint(2)},
		Account: "Bank of Ireland",
		Amount:  1234.50,
		Date:    time.Date(2024, 1, 15, 9, 0, 0, 0, time.UTC),
		Vendor:  "Work",
	}

	rows := s.mock.NewRows([]string{"id", "account", "amount", "date", "vendor"}).
		AddRow(e1.ID, e1.Account, e1.Amount, e1.Date, e1.Vendor).
		AddRow(e2.ID, e2.Account, e2.Amount, e2.Date, e2.Vendor)
	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	expenses, err := s.repo.GetExpenses()
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), expenses)
	assert.Equal(s.T(), e1.ID, expenses[0].ID, "Income 1 ID matches expected")
	assert.Equal(s.T(), e1.Account, expenses[0].Account, "Income 1 Account matches expected")
	assert.Equal(s.T(), e1.Amount, expenses[0].Amount, "Income 1 Amount matches expected")
	assert.Equal(s.T(), e1.Date, expenses[0].Date, "Income 1 Date matches expected")
	assert.Equal(s.T(), e1.Vendor, expenses[0].Vendor, "Income 1 Vendor matches expected")
	assert.Equal(s.T(), e2.ID, expenses[1].ID, "Income 2 ID matches expected")
	assert.Equal(s.T(), e2.Account, expenses[1].Account, "Income 2 Account matches expected")
	assert.Equal(s.T(), e2.Amount, expenses[1].Amount, "Income 2 Amount matches expected")
	assert.Equal(s.T(), e2.Date, expenses[1].Date, "Income 2 Date matches expected")
	assert.Equal(s.T(), e2.Vendor, expenses[1].Vendor, "Income 2 Vendor matches expected")
}
