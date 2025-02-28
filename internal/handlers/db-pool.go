package handlers

import (
	"net/http"

	"github.com/FDS-Studio/db-gateway/internal/services"
	"github.com/gin-gonic/gin"
)

type DbConnectionPoolHandler struct {
	dbConnectionPoolService *services.DbConnectionPoolService
}

func NewDbConnectionPoolHandler(dbConnectionPoolService *services.DbConnectionPoolService) *DbConnectionPoolHandler {
	return &DbConnectionPoolHandler{
		dbConnectionPoolService: dbConnectionPoolService,
	}
}

// ListDbConnectionPoolNames godoc
// @Summary Get a list of db connection pool
// @Description Get a list of all db connection pool
// @Produce application/json
// @Tags DbConnPool
// @Success 200 {array} []string
// @Failure 500 {object} map[string]string
// @Router /db-pool/all [get]
func (dbcph *DbConnectionPoolHandler) ListDbConnectionPoolNames(c *gin.Context) {
	result, err := dbcph.dbConnectionPoolService.ListDbConnectionPoolNames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
