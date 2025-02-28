package routes

import (
	"github.com/FDS-Studio/db-gateway/internal/handlers"
	"github.com/gin-gonic/gin"
)

func DbConfigRoutes(rg *gin.RouterGroup, dbch *handlers.DbConfigHandler) {
	rg.GET("/all", dbch.ListDBConfigsHandler)
	rg.POST("/", dbch.CreateDBConfigHandler)
	rg.PUT("/", dbch.UpdateDBConfigHandler)
	rg.DELETE("/:name", dbch.DeleteDBConfigHandler)
}
