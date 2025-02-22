package configuration

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigInfo struct {
	Server   Server   `yaml:"server"`
	DbServer DbServer `yaml:"db-server"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DbServer struct {
	Host          string `yaml:"host"`
	Port          string `yaml:"port"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
	DefaultDbName string `yaml:"defaultDbName"`
}

var Config *ConfigInfo

func LoadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&Config); err != nil {
		return err
	}

	return nil
}
