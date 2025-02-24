package dbpoll

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/FDS-Studio/db-gateway/internal/config"
	_ "github.com/lib/pq"
)

type DbConnectionPool struct {
	dbConnections map[string]*sql.DB
}

func New() *DbConnectionPool {
	dbcp := &DbConnectionPool{
		dbConnections: make(map[string]*sql.DB),
	}

	return dbcp
}

func (dbcp *DbConnectionPool) Connect(dbConfig config.DbConfig) error {
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

func (dbcp *DbConnectionPool) Get(name string) (*sql.DB, error) {
	val, ok := dbcp.dbConnections[name]
	if ok {
		return val, nil
	}

	return nil, errors.New("config not found for name: " + name)
}

func (dbcp *DbConnectionPool) Close(name string) error {
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

func (dbcp *DbConnectionPool) CloseAll() error {
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

func (dbcp *DbConnectionPool) CheckStatus(name string) bool {
	_, ok := dbcp.dbConnections[name]
	return ok
}
