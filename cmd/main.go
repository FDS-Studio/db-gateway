package main

import (
	"fmt"

	docs "github.com/FDS-Studio/db-gateway/docs"
	"github.com/FDS-Studio/db-gateway/internal/configuration"
	"github.com/FDS-Studio/db-gateway/internal/controllers"
	"github.com/FDS-Studio/db-gateway/internal/repositories"
	"github.com/FDS-Studio/db-gateway/internal/routes"
	"github.com/FDS-Studio/db-gateway/internal/services"
	"github.com/gin-gonic/gin"

	dbconnservice "github.com/FDS-Studio/db-gateway/internal/services/db-conn-service"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title FDS Studio DB GATEWAY
// @version 1.0
// @description ...

// @host localhost:8080
// @BasePath /api/v1
func main() {
	err := configuration.LoadConfig("configs/config.yaml")
	if err != nil {
		panic(err)
	}

	dbConn, err := dbconnservice.New()
	if err != nil {
		panic(err)
	}
	defer dbConn.CloseAll()

	pgRep := repositories.NewDbRepository()
	pgService := services.NewDbService(pgRep, dbConn)
	pgController := controllers.NewDbController(pgService)

	databases, err := pgService.GetAllDb()
	if err != nil {
		panic(err)
	}

	for _, v := range databases {
		if err := dbConn.Connect(v.Name); err != nil {
			fmt.Printf("Ошибка подключения к базе данных %s: %s\n", v.Name, err)
		}
	}

	fmt.Println(dbConn)

	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		routes.DbRoutes(v1.Group("/database"), pgController)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	if err := r.Run(fmt.Sprintf(":%v", configuration.Config.Server.Port)); err != nil {
		panic(err)
	}
}
