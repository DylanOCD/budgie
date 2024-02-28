/*
 * Copyright (c) 2024 Dylan O' Connor Desmond
 */

package handlers

import (
	"net/http"

	"github.com/DylanOCD/budgie/backend/internal/config"
	"github.com/DylanOCD/budgie/backend/internal/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repository *repository.Repository
	conf       *config.Conf
}

func New(repo *repository.Repository, conf *config.Conf) Handler {
	return Handler{repo, conf}
}

func (handler Handler) Ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Tweet tweet!"})
}
