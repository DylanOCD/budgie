/*
 * Copyright (c) 2024 Dylan O' Connor Desmond
 */

package database

import (
	"fmt"
	"log"

	"github.com/DylanOCD/budgie/backend/config"
	"github.com/DylanOCD/budgie/backend/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(conf config.Conf) *gorm.DB {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		conf.DatabaseUsername,
		conf.DatabasePassword,
		conf.DatabaseHost,
		conf.DatabasePort,
		conf.DatabaseName,
	)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	seeded := db.Migrator().HasTable(&models.Income{})
	db.AutoMigrate(&models.Income{}, &models.Expense{})
	if !seeded {
		Seed(db)
	}

	return db
}
