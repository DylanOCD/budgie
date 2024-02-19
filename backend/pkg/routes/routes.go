/*
 * Copyright (c) 2024 Dylan O' Connor Desmond
 */

package routes

import (
	"github.com/DylanOCD/budgie/backend/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine, handler handlers.Handler) {
	v1 := router.Group("budgie/v1")
	{
		// Ping
		v1.GET("/ping", handler.Ping)

		// Income
		v1.GET("/income", handler.GetIncomes)

		// Expense
		v1.GET("/expense", handler.GetExpenses)
	}
}
