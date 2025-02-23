package dbpoll

import (
	"database/sql"
	"errors"
	"fmt"

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

func (dbcp *DbConnectionPool) Connect(dbConfig config.Database) error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
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

func (dbc *DbConnectionPool) Get(name string) (*sql.DB, error) {
	val, ok := dbc.dbConnections[name]
	if ok {
		return val, nil
	}

	return nil, errors.New("value not found for name: " + name)
}

func (dbc *DbConnectionPool) Close(name string) error {
	db, ok := dbc.dbConnections[name]
	if !ok {
		return errors.New("no connection found for name: " + name)
	}
	err := db.Close()
	if err != nil {
		return err
	}
	delete(dbc.dbConnections, name)
	return nil
}

func (dbc *DbConnectionPool) CloseAll() {
	for k := range dbc.dbConnections {
		dbc.Close(k)
	}
}
