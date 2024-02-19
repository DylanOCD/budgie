/*
 * Copyright (c) 2024 Dylan O' Connor Desmond
 */

package models

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	Account string    `json:"account"`
	Amount  float64   `json:"amount"`
	Date    time.Time `json:"date"`
	Vendor  string    `json:"vendor"`
}
