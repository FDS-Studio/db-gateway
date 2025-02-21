package controllers

import (
	"net/http"

	"github.com/FDS-Studio/db-gateway/internal/models"
	"github.com/FDS-Studio/db-gateway/internal/services"
	"github.com/gin-gonic/gin"
)

type DbController struct {
	dbService *services.DbService
}

func NewDbController(dbService *services.DbService) *DbController {
	return &DbController{
		dbService: dbService,
	}
}

func (dbc *DbController) CreateDb(c *gin.Context) {
	var database models.Database
	err := c.ShouldBindJSON(&database)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	createdDb, err := dbc.dbService.CreateDb(database)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdDb)
}
