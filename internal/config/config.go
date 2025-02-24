package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type DbConfig struct {
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

const DbConfigPath string = "configs/db"
const serverPath string = "configs/server.yaml"

func LoadDbConfig() ([]DbConfig, error) {
	configs, err := os.ReadDir(DbConfigPath)
	if err != nil {
		return nil, err
	}

	databases := make([]DbConfig, len(configs))

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

func readDbConfig(name string) (DbConfig, error) {
	filePath := filepath.Join(DbConfigPath, name)
	file, err := os.Open(filePath)
	if err != nil {
		return DbConfig{}, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	database := DbConfig{}

	err = d.Decode(&database)
	if err != nil {
		return DbConfig{}, err
	}

	return database, nil
}
