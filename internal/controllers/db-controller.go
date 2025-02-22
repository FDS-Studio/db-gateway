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

// CreateDatabase godoc
// @Summary Create a database
// @Description Create a new database with the given name
// @Param database body models.Database true "Database name"
// @Produce application/json
// @Tags Database
// @Success         200 {object} map[string]string
// @Failure         400 {object} map[string]string
// @Router /database/ [post]
func (dbc *DbController) CreateDb(c *gin.Context) {
	var database models.Database
	if err := c.ShouldBindJSON(&database); err != nil || database.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request, database name is required"})
		return
	}

	result, err := dbc.dbService.CreateDb(database)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (dbc *DbController) GetAllDb(c *gin.Context) {
	result, err := dbc.dbService.GetAllDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

// DeleteDatabase godoc
// @Summary Delete a database
// @Description Delete a new database with the given name
// @Param database body models.Database true "Database name"
// @Produce application/json
// @Tags Database
// @Success         200 {object} map[string]string
// @Failure         400 {object} map[string]string
// @Router /database/ [delete]
func (dbc *DbController) DropDb(c *gin.Context) {
	var database models.Database
	if err := c.ShouldBindJSON(&database); err != nil || database.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request, database name is required"})
		return
	}

	result, err := dbc.dbService.DropDb(database)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}
