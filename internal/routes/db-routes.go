package routes

import (
	"github.com/FDS-Studio/db-gateway/internal/controllers"
	"github.com/gin-gonic/gin"
)

func DbRoutes(rg *gin.RouterGroup, dbc *controllers.DbController) {
	rg.POST("/", dbc.CreateDb)
	rg.GET("/all", dbc.GetAllDb)
	rg.DELETE("/", dbc.DropDb)
}
