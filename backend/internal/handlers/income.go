/*
 * Copyright (c) 2024 Dylan O' Connor Desmond
 */

package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler Handler) GetIncomes(c *gin.Context) {
	incomes, err := handler.repository.GetIncomes()
	if err != nil {
		message := fmt.Sprintf("Failed to get incomes: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}
	c.JSON(http.StatusOK, incomes)
}
