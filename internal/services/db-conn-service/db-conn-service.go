package dbconnservice

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/FDS-Studio/db-gateway/internal/configuration"
	_ "github.com/lib/pq"
)

type DbConnections struct {
	dbConnections map[string]*sql.DB
	conStr        string
}

func New() (*DbConnections, error) {
	conStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=",
		configuration.Config.DbServer.Host, configuration.Config.DbServer.Port,
		configuration.Config.DbServer.Username, configuration.Config.DbServer.Password)

	dbc := &DbConnections{
		dbConnections: make(map[string]*sql.DB),
		conStr:        conStr,
	}

	defaultDbName := configuration.Config.DbServer.DefaultDbName

	if defaultDbName != "" {
		if err := dbc.Connect(defaultDbName); err != nil {
			return nil, fmt.Errorf("failed to connect to default database: %w", err)
		}
	}

	return dbc, nil
}

func (dbc *DbConnections) Get(name string) (*sql.DB, error) {
	val, ok := dbc.dbConnections[name]
	if ok {
		return val, nil
	}

	return nil, errors.New("value not found for name: " + name)
}

func (dbc *DbConnections) Connect(name string) error {
	_, ok := dbc.dbConnections[name]
	if ok {
		return nil
	}

	db, err := sql.Open("postgres", dbc.conStr+name)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	dbc.dbConnections[name] = db
	return nil
}

func (dbc *DbConnections) Close(name string) error {
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

func (dbc *DbConnections) CloseAll() {
	for k := range dbc.dbConnections {
		dbc.Close(k)
	}
}
