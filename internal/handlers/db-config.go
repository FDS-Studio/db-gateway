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

// CreateDbConfig godoc
// @Summary Create a db config
// @Description Create a new db config
// @Param dbConfig body models.DbConfig true "Database configuration with host, port, username, password and name"
// @Produce application/json
// @Tags DbConfig
// @Success         200 {object} map[string]string
// @Failure         400 {object} map[string]string
// @Router /db-configs/ [post]
func (dbch *DbConfigHandler) CreateDBConfigHandler(c *gin.Context) {
	var dbConfig models.DbConfig
	if err := c.ShouldBindJSON(&dbConfig); err != nil || dbConfig.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request, database config name is required"})
		return
	}

	err := dbch.dbConfigService.CreateDBConfigHandler(dbConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Database config created"})
}

func (dbch *DbConfigHandler) UpdateDBConfigHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Database config updated"})
}

func (dbch *DbConfigHandler) GetDBConfigHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Fetched database config"})
}

// DeleteDBConfigHandler godoc
// @Summary Delete a db config
// @Description Delete a database configuration by its name
// @Produce application/json
// @Param name path string true "Name of the database config to delete"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /db-configs/{name} [delete]
func (dbch *DbConfigHandler) DeleteDBConfigHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request, database config name is required"})
		return
	}

	err := dbch.dbConfigService.DeleteDBConfigHandler(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Database config deleted"})
}

// ListDBConfigsHandler godoc
// @Summary Get a list of db configs
// @Description Get a list of all db configs and their statuses
// @Produce application/json
// @Tags DbConfig
// @Success 200 {array} models.DbConfig
// @Failure 500 {object} map[string]string
// @Router /db-configs/all [get]
func (dbch *DbConfigHandler) ListDBConfigsHandler(c *gin.Context) {
	result, err := dbch.dbConfigService.ListDBConfigsHandler()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
