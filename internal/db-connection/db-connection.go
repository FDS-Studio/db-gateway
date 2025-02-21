package dbconnection

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/FDS-Studio/db-gateway/internal/configuration"
	_ "github.com/lib/pq"
)

var dbConnections map[string]*sql.DB
var conStr string

func Run() {
	conStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=",
		configuration.Config.DbServer.Host, configuration.Config.DbServer.Port,
		configuration.Config.DbServer.Username, configuration.Config.DbServer.Password)
}

func Get(name string) (*sql.DB, error) {
	val, ok := dbConnections[name]
	if ok {
		return val, nil
	}

	return nil, errors.New("value not found for name: " + name)
}

func Connect(name string) error {
	db, err := sql.Open("postgres", conStr+name)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	dbConnections[name] = db
	return nil
}

func Close(name string) error {
	db, ok := dbConnections[name]
	if !ok {
		return errors.New("no connection found for name: " + name)
	}
	err := db.Close()
	if err != nil {
		return err
	}
	delete(dbConnections, name)
	return nil
}
