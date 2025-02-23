package handlers

import (
	"net/http"

	"github.com/FDS-Studio/db-gateway/internal/models"
	"github.com/FDS-Studio/db-gateway/internal/services"
	"github.com/gin-gonic/gin"
)

type DbConfigHandler struct {
	dbConfigService *services.DbConfigService
}

func NewDbConfigHandler(dbConfigService *services.DbConfigService) *DbConfigHandler {
	return &DbConfigHandler{
		dbConfigService: dbConfigService,
	}
}

// CreateDatabase godoc
// @Summary Create a database
// @Description Create a new database with the given name
// @Param database body models.DbConfig true "Database name"
// @Produce application/json
// @Tags DbConfig
// @Success         200 {object} map[string]string
// @Failure         400 {object} map[string]string
// @Router /database/ [post]
func (dbch *DbConfigHandler) CreateDBConfigHandler(c *gin.Context) {
	var dbConfig models.DbConfig
	if err := c.ShouldBindJSON(&dbConfig); err != nil || dbConfig.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request, database name is required"})
		return
	}

	err := dbch.dbConfigService.CreateDBConfigHandler(dbConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Database config created"})
}

func UpdateDBConfigHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Database config updated"})
}

func GetDBConfigHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Fetched database config"})
}

func DeleteDBConfigHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Database config deleted"})
}

func ListDBConfigsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List of database configs"})
}
