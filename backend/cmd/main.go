/*
 * Copyright (c) 2024 Dylan O' Connor Desmond
 */

package main

import (
	"fmt"
	"log"

	"github.com/DylanOCD/budgie/backend/config"
	"github.com/DylanOCD/budgie/backend/pkg/database"
	"github.com/DylanOCD/budgie/backend/pkg/handlers"
	"github.com/DylanOCD/budgie/backend/pkg/repository"
	"github.com/DylanOCD/budgie/backend/pkg/routes"
	"github.com/DylanOCD/budgie/backend/rootdir"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB, handler handlers.Handler) *gin.Engine {
	router := gin.Default()
	routes.AddRoutes(router, handler)

	return router
}

func main() {
	conf, err := config.Load(rootdir.Root + "/config")
	if err != nil {
		message := fmt.Sprintf("Failed to load conf: %v", err)
		log.Fatal(message)
	}

	db := database.Connect(conf)
	r := repository.New(db)
	handler := handlers.New(&r, &conf)
	router := setupRouter(db, handler)

	// Listen and Server in 0.0.0.0:8080
	err = router.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting router: %v", err)
	}
}
