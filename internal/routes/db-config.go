package routes

import (
	"github.com/FDS-Studio/db-gateway/internal/handlers"
	"github.com/gin-gonic/gin"
)

func DbRoutes(rg *gin.RouterGroup, dbch *handlers.DbConfigHandler) {
	rg.POST("/", dbch.CreateDBConfigHandler)
}
