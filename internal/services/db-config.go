package services

import (
	"fmt"
	"os"
	"path"

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
		return err
	}

	filePath := path.Join(config.DbConfigPath, dbConfig.Name+".yaml")

	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		return fmt.Errorf("file %s already exists", filePath)
	}

	file, err := os.Create(filePath)
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
