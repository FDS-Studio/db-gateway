package services

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/FDS-Studio/db-gateway/internal/config"
	"github.com/FDS-Studio/db-gateway/internal/models"
	dbpoll "github.com/FDS-Studio/db-gateway/internal/services/db-poll"
	"gopkg.in/yaml.v2"
)

type DbConfigService struct {
	dbConnectionPool *dbpoll.DbConnectionPool
}

func NewDbConfigService(dbConnectionPool *dbpoll.DbConnectionPool) *DbConfigService {
	return &DbConfigService{
		dbConnectionPool: dbConnectionPool,
	}
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

func (dbcs *DbConfigService) UpdateDBConfigHandler() {
}

func (dbcs *DbConfigService) GetDBConfigHandler() {
}

func (dbcs *DbConfigService) DeleteDBConfigHandler(name string) error {
	filePath := path.Join(config.DbConfigPath, name+".yaml")

	ok := dbcs.dbConnectionPool.CheckStatus(name)
	if ok {
		return errors.New("config in progress")
	}

	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}

func (dbcs *DbConfigService) ListDBConfigsHandler() ([]models.DbConfig, error) {
	files, err := os.ReadDir(config.DbConfigPath)
	if err != nil {
		return nil, err
	}

	dbConfigs := make([]models.DbConfig, 0)

	for _, file := range files {
		dbConfig, err := dbcs.readDbConfig(file.Name())
		if err != nil {
			return nil, err
		}

		ok := dbcs.dbConnectionPool.CheckStatus(dbConfig.Name)

		dbConfig.IsRun = ok
		dbConfigs = append(dbConfigs, dbConfig)
	}

	return dbConfigs, nil
}

func (dbcs *DbConfigService) readDbConfig(fileName string) (models.DbConfig, error) {
	filePath := path.Join(config.DbConfigPath, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		return models.DbConfig{}, err
	}
	defer file.Close()

	var dbConfig models.DbConfig
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&dbConfig); err != nil {
		return models.DbConfig{}, err
	}

	return dbConfig, nil
}
