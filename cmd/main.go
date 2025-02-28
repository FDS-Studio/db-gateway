package main

import (
	"github.com/FDS-Studio/db-gateway/docs"
	"github.com/FDS-Studio/db-gateway/internal/config"
	"github.com/FDS-Studio/db-gateway/internal/handlers"
	"github.com/FDS-Studio/db-gateway/internal/routes"
	"github.com/FDS-Studio/db-gateway/internal/services"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title FDS Studio DB GATEWAY
// @version 1.0
// @description ...

// @host localhost:8080
// @BasePath /api/v1
func main() {
	databases, err := config.LoadDbConfig()
	if err != nil {
		panic(err)
	}

	serverConfig, err := config.LoadServerConfig()
	if err != nil {
		panic(err)
	}

	dbConnPollService := services.NewDbConnectionPoolService()
	dbConnPollHandler := handlers.NewDbConnectionPoolHandler(dbConnPollService)
	dbConfgService := services.NewDbConfigService(dbConnPollService)
	dbConfigHandler := handlers.NewDbConfigHandler(dbConfgService)

	for _, dbConfig := range databases {
		if dbConfig.AutoRun {
			err := dbConnPollService.Connect(dbConfig)
			if err != nil {
				panic(err)
			}
		}
	}

	defer dbConnPollService.CloseAll()

	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		routes.DbConfigRoutes(v1.Group("/db-configs"), dbConfigHandler)
		routes.DbConnectionPoolRoutes(v1.Group("/db-pool"), dbConnPollHandler)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	if err := r.Run(serverConfig.Address); err != nil {
		panic(err)
	}
}
