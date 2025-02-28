package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/FDS-Studio/db-gateway/internal/config"
	_ "github.com/lib/pq"
)

type DbConnectionPoolService struct {
	dbConnections map[string]*sql.DB
}

func NewDbConnectionPoolService() *DbConnectionPoolService {
	dbcp := &DbConnectionPoolService{
		dbConnections: make(map[string]*sql.DB),
	}

	return dbcp
}

func (dbcps *DbConnectionPoolService) ListDbConnectionPoolNames() ([]string, error) {
	if len(dbcps.dbConnections) <= 0 {
		return nil, errors.New("no database connections available")
	}

	keys := make([]string, 0, len(dbcps.dbConnections))

	for key := range dbcps.dbConnections {
		keys = append(keys, key)
	}

	return keys, nil
}

func (dbcp *DbConnectionPoolService) Connect(dbConfig config.DbConfig) error {
	connStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Name)

	_, ok := dbcp.dbConnections[dbConfig.Name]
	if ok {
		return nil
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	dbcp.dbConnections[dbConfig.Name] = db
	return nil
}

func (dbcp *DbConnectionPoolService) Get(name string) (*sql.DB, error) {
	val, ok := dbcp.dbConnections[name]
	if ok {
		return val, nil
	}

	return nil, errors.New("config not found for name: " + name)
}

func (dbcp *DbConnectionPoolService) Close(name string) error {
	db, ok := dbcp.dbConnections[name]
	if !ok {
		return errors.New("no connection found for name: " + name)
	}
	err := db.Close()
	if err != nil {
		return err
	}
	delete(dbcp.dbConnections, name)
	return nil
}

func (dbcp *DbConnectionPoolService) CloseAll() error {
	if len(dbcp.dbConnections) == 0 {
		return errors.New("no connections to close")
	}

	for k, db := range dbcp.dbConnections {
		if err := db.Close(); err != nil {
			log.Printf("Error when closing connection with %s: %v", k, err)
			continue
		}
		delete(dbcp.dbConnections, k)
		log.Printf("Connection to %s closed successfully", k)
	}

	return nil
}

func (dbcp *DbConnectionPoolService) CheckStatus(name string) bool {
	_, ok := dbcp.dbConnections[name]
	return ok
}
