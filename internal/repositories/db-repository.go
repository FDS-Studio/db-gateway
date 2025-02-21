package repositories

import (
	"database/sql"

	"github.com/FDS-Studio/db-gateway/internal/models"
)

type DbRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *DbRepository {
	return &DbRepository{
		db: db,
	}
}

func (dbr *DbRepository) CreateDb(database models.Database) (*models.Database, error) {
	_, err := dbr.db.Exec("CREATE DATABASE $1", database.Name)
	if err != nil {
		return nil, err
	}

	return &database, nil
}

func (dbr *DbRepository) DropDb() {

}

func (dbr *DbRepository) GetAllDb() {

}

func (dbr *DbRepository) GetDbByName(nameDb string) {

}
