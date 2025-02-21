package main

import (
	"github.com/FDS-Studio/db-gateway/internal/configuration"
	dbconnection "github.com/FDS-Studio/db-gateway/internal/db-connection"
)

func main() {
	err := configuration.LoadConfig("configs/config.yaml")
	if err != nil {
		panic(err)
	}
	dbconnection.Run()
}
