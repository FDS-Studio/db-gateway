package main

import (
	"github.com/FDS-Studio/db-gateway/docs"
	"github.com/FDS-Studio/db-gateway/internal/config"
	"github.com/FDS-Studio/db-gateway/internal/handlers"
	"github.com/FDS-Studio/db-gateway/internal/routes"
	"github.com/FDS-Studio/db-gateway/internal/services"
	dbpoll "github.com/FDS-Studio/db-gateway/internal/services/db-poll"
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

	dbConnPoll := dbpoll.New()

	for _, dbConfig := range databases {
		err := dbConnPoll.Connect(dbConfig)
		if err != nil {
			panic(err)
		}
	}

	defer dbConnPoll.CloseAll()

	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		dbConfgService := services.NewDbConfigService(dbConnPoll)
		dbConfigHandler := handlers.NewDbConfigHandler(dbConfgService)
		routes.DbRoutes(v1.Group("/db-configs"), dbConfigHandler)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	if err := r.Run(serverConfig.Address); err != nil {
		panic(err)
	}
}
