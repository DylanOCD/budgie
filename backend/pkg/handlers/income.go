/*
 * Copyright (c) 2024 Dylan O' Connor Desmond
 */

package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler Handler) GetIncomes(context *gin.Context) {
	incomes, err := handler.repository.GetIncomes()
	if err != nil {
		message := fmt.Sprintf("Failed to get incomes: %v", err)
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}
	context.IndentedJSON(http.StatusOK, incomes)
}
