package services

import (
	"github.com/FDS-Studio/db-gateway/internal/models"
	"github.com/FDS-Studio/db-gateway/internal/repositories"
)

type DbService struct {
	dbRepository *repositories.DbRepository
}

func NewDbService(dbRepository *repositories.DbRepository) *DbService {
	return &DbService{
		dbRepository: dbRepository,
	}
}

func (dbs *DbService) CreateDb(database models.Database) (*models.Database, error) {
	resp, err := dbs.dbRepository.CreateDb(database)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
