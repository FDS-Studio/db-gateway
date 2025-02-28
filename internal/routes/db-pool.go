package routes

import (
	"github.com/FDS-Studio/db-gateway/internal/handlers"
	"github.com/gin-gonic/gin"
)

func DbConnectionPoolRoutes(rg *gin.RouterGroup, dbch *handlers.DbConnectionPoolHandler) {
	rg.GET("/all", dbch.ListDbConnectionPoolNames)
}
