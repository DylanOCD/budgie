/*
 * Copyright (c) 2024 Dylan O' Connor Desmond
 */

package handlers

import (
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DylanOCD/budgie/backend/internal/config"
	"github.com/DylanOCD/budgie/backend/internal/models"
	"github.com/DylanOCD/budgie/backend/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IncomeSuite struct {
	suite.Suite
	db      *gorm.DB
	mock    sqlmock.Sqlmock
	repo    repository.IRepository
	handler Handler
	router  *gin.Engine
}

func (s *IncomeSuite) SetupTest() {
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

	r := repository.New(s.db)
	s.repo = &r
	conf := config.Conf{}
	h := New(&r, &conf)
	s.handler = h
	router := gin.Default()
	s.router = router
	AddRoutes(s.router, s.handler)
}

func (s *IncomeSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(IncomeSuite))
}

func (s *IncomeSuite) TestGetIncomes() {
	query := `SELECT * FROM "incomes" WHERE "incomes"."deleted_at" IS NULL`
	var incomes = []models.Income{
		{
			Model:   gorm.Model{ID: uint(1)},
			Account: "Bank of Ireland",
			Amount:  1234.50,
			Date:    time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC),
			Source:  "Work",
		},
		{
			Model:   gorm.Model{ID: uint(2)},
			Account: "Bank of Ireland",
			Amount:  1234.50,
			Date:    time.Date(2024, 1, 15, 9, 0, 0, 0, time.UTC),
			Source:  "Work",
		},
	}

	tests := []struct {
		LoadDatabase   []models.Income
		HttpStatusCode int
		ExpectedBody   string
	}{
		{
			LoadDatabase:   incomes,
			HttpStatusCode: http.StatusOK,
			ExpectedBody: `[
	{
		"ID": 1,
		"CreatedAt": "0001-01-01T00:00:00Z",
		"UpdatedAt": "0001-01-01T00:00:00Z",
		"DeletedAt": null,
		"account": "Bank of Ireland",
		"amount": 1234.5,
		"date": "2024-01-01T09:00:00Z",
		"source": "Work"
	},
	{
		"ID": 2,
		"CreatedAt": "0001-01-01T00:00:00Z",
		"UpdatedAt": "0001-01-01T00:00:00Z",
		"DeletedAt": null,
		"account": "Bank of Ireland",
		"amount": 1234.5,
		"date": "2024-01-15T09:00:00Z",
		"source": "Work"
	}
]`,
		},
		/*
			{
				LoadDatabase:   []models.Income{},
				HttpStatusCode: http.StatusBadRequest,
				ExpectedBody:   `{"message": "Failed to get incomes: err"}`,
			},
		*/
	}

	for _, test := range tests {
		rows := s.mock.NewRows([]string{
			"id",
			"account",
			"amount",
			"date",
			"source",
		})
		for _, income := range incomes {
			rows.AddRow(
				income.ID,
				income.Account,
				income.Amount,
				income.Date,
				income.Source,
			)
		}
		s.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/budgie/v1/income", nil)
		s.router.ServeHTTP(w, req)
		responseBody, err := io.ReadAll(w.Body)
		if err != nil {
			s.T().Error("failed to read response body")
		}

		assert.Equal(s.T(), test.HttpStatusCode, w.Code)
		assert.Equal(s.T(), test.ExpectedBody, string(responseBody))
	}
}
