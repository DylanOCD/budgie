/*
 * Copyright (c) 2024 Dylan O' Connor Desmond
 */

package repository

import (
	"github.com/DylanOCD/budgie/backend/pkg/models"
	"gorm.io/gorm"
)

type IRepository interface {
	// Income
	GetIncomes() ([]*models.Income, error)
	GetExpenses() ([]*models.Expense, error)
}

type Repository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Repository {
	return Repository{db}
}

func (repository *Repository) GetIncomes() ([]*models.Income, error) {
	incomes := make([]*models.Income, 0)
	err := repository.DB.Find(&incomes).Error
	return incomes, err
}

func (repository *Repository) GetExpenses() ([]*models.Expense, error) {
	expenses := make([]*models.Expense, 0)
	err := repository.DB.Find(&expenses).Error
	return expenses, err
}
