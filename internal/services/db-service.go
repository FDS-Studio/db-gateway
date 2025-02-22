package services

import (
	"github.com/FDS-Studio/db-gateway/internal/configuration"
	"github.com/FDS-Studio/db-gateway/internal/models"
	"github.com/FDS-Studio/db-gateway/internal/repositories"
	dbconnservice "github.com/FDS-Studio/db-gateway/internal/services/db-conn-service"
)

type DbService struct {
	dbRepository  *repositories.DbRepository
	dbConnections *dbconnservice.DbConnections
}

func NewDbService(dbRepository *repositories.DbRepository, dbConnection *dbconnservice.DbConnections) *DbService {
	return &DbService{
		dbRepository:  dbRepository,
		dbConnections: dbConnection,
	}
}

func (dbs *DbService) CreateDb(database models.Database) (models.Database, error) {
	db, err := dbs.dbConnections.Get(configuration.Config.DbServer.DefaultDbName)
	if err != nil {
		return database, err
	}

	result, err := dbs.dbRepository.CreateDb(db, database)
	if err != nil {
		return database, err
	}

	return result, nil
}

func (dbs *DbService) DropDb(database models.Database) (models.Database, error) {
	db, err := dbs.dbConnections.Get(configuration.Config.DbServer.DefaultDbName)
	if err != nil {
		return database, err
	}

	result, err := dbs.dbRepository.DropDb(db, database)
	if err != nil {
		return database, err
	}

	return result, nil
}

func (dbs *DbService) GetAllDb() ([]models.Database, error) {
	db, err := dbs.dbConnections.Get(configuration.Config.DbServer.DefaultDbName)
	if err != nil {
		return nil, err
	}

	result, err := dbs.dbRepository.GetAllDb(db)
	if err != nil {
		return nil, err
	}

	return result, nil
}
