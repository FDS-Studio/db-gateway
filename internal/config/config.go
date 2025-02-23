package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Server struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Address string
}

const DbPath string = "configs/db"
const serverPath string = "configs/server.yaml"

func LoadDbConfig() ([]Database, error) {
	configs, err := os.ReadDir(DbPath)
	if err != nil {
		return nil, err
	}

	databases := make([]Database, len(configs))

	for i := range configs {
		database, err := readDbConfig(configs[i].Name())
		if err != nil {
			return nil, err
		}
		databases[i] = database
	}

	return databases, nil
}

func LoadServerConfig() (Server, error) {
	file, err := os.Open(serverPath)
	if err != nil {
		return Server{}, err
	}
	defer file.Close()

	var server Server
	d := yaml.NewDecoder(file)

	if err := d.Decode(&server); err != nil {
		return Server{}, err
	}

	server.Address = fmt.Sprintf("%v:%v", server.Host, server.Port)

	return server, nil
}

func readDbConfig(name string) (Database, error) {
	filePath := filepath.Join(DbPath, name)
	file, err := os.Open(filePath)
	if err != nil {
		return Database{}, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	database := Database{}

	err = d.Decode(&database)
	if err != nil {
		return Database{}, err
	}

	return database, nil
}
