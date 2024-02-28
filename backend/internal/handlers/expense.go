/*
 * Copyright (c) 2024 Dylan O' Connor Desmond
 */

package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler Handler) GetExpenses(c *gin.Context) {
	expenses, err := handler.repository.GetExpenses()
	if err != nil {
		message := fmt.Sprintf("Failed to get expenses: %v", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}
	c.IndentedJSON(http.StatusOK, expenses)
}
