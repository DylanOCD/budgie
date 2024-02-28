/*
 * Copyright (c) 2024 Dylan O' Connor Desmond
 */

package database

import (
	"fmt"
	"time"

	"github.com/DylanOCD/budgie/backend/internal/models"
	"gorm.io/gorm"
)

var incomes = []models.Income{
	{
		Account: "Bank of Ireland",
		Amount:  1234.50,
		Date:    time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC),
		Source:  "Work",
	},
	{
		Account: "Bank of Ireland",
		Amount:  1234.50,
		Date:    time.Date(2024, 1, 15, 9, 0, 0, 0, time.UTC),
		Source:  "Work",
	},
}

var expenses = []models.Expense{
	{
		Account: "Bank of Ireland",
		Amount:  12.50,
		Date:    time.Date(2024, 1, 7, 18, 32, 15, 0, time.UTC),
		Vendor:  "Tesco",
	},
	{
		Account: "Bank of Ireland",
		Amount:  20.00,
		Date:    time.Date(2024, 1, 22, 19, 45, 0, 0, time.UTC),
		Vendor:  "Omniplex",
	},
}

func Seed(database *gorm.DB) {
	for _, income := range incomes {
		if result := database.Create(&income); result.Error != nil {
			fmt.Printf(
				"Failed to create income with Account %s, Amount %f, Date %v, Source %s: %v",
				income.Account,
				income.Amount,
				income.Date.String(),
				income.Source,
				result.Error,
			)
		}
	}

	for _, expense := range expenses {
		if result := database.Create(&expense); result.Error != nil {
			fmt.Printf(
				"Failed to create expense with Account %s, Amount %f, Date %v, Vendor %s: %v",
				expense.Account,
				expense.Amount,
				expense.Date.String(),
				expense.Vendor,
				result.Error,
			)
		}
	}
}
