package repositories

import (
	"database/sql"
	"fmt"

	"github.com/FDS-Studio/db-gateway/internal/models"
	"github.com/lib/pq"
)

type DbRepository struct {
}

func NewDbRepository() *DbRepository {
	return &DbRepository{}
}

func (dbr *DbRepository) GetAllDb(db *sql.DB) ([]models.Database, error) {
	rows, err := db.Query("SELECT datname AS name FROM pg_database WHERE datistemplate = false;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var databases []models.Database
	for rows.Next() {
		var database models.Database
		if err := rows.Scan(&database.Name); err != nil {
			return nil, err
		}
		databases = append(databases, database)
	}

	return databases, nil
}

func (dbr *DbRepository) CreateDb(db *sql.DB, database models.Database) (models.Database, error) {
	dbName := pq.QuoteIdentifier(database.Name)
	query := fmt.Sprintf("CREATE DATABASE %s", dbName)

	_, err := db.Exec(query)
	if err != nil {
		return models.Database{}, err
	}

	return database, nil
}

func (dbr *DbRepository) DropDb(db *sql.DB, database models.Database) (models.Database, error) {
	dbName := pq.QuoteIdentifier(database.Name)
	query := fmt.Sprintf("DROP DATABASE %s", dbName)

	_, err := db.Exec(query)
	if err != nil {
		return models.Database{}, err
	}

	return database, nil
}
