package services

import (
	"fmt"
	"os"

	"github.com/FDS-Studio/db-gateway/internal/config"
	"github.com/FDS-Studio/db-gateway/internal/models"
	"gopkg.in/yaml.v2"
)

type DbConfigService struct {
}

func NewDbConfigService() *DbConfigService {
	return &DbConfigService{}
}

func (*DbConfigService) CreateDBConfigHandler(dbConfig models.DbConfig) error {
	yamlData, err := yaml.Marshal(&dbConfig)
	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}

	file, err := os.Create(config.DbPath + "/" + dbConfig.Name + ".yaml")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(yamlData)
	if err != nil {
		return err
	}

	return nil
}

func UpdateDBConfigHandler() {
}

func GetDBConfigHandler() {
}

func DeleteDBConfigHandler() {
}

func ListDBConfigsHandler() {
}
