package configs

import (
	"gopkg.in/yaml.v2"
	"os"
)

type AppConfig struct {
	Database *Database `yaml:"db"`
	Nats     *Nats     `yaml:"nats"`
	Server   *Server   `yaml:"server"`
}

type Database struct {
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Nats struct {
	ClusterID string `yaml:"clusterID"`
	ClientID  string `yaml:"clientID"`
	Subject   string `yaml:"subject"`
}

type Server struct {
	Port string `yaml:"port"`
}

func Initialize() (*AppConfig, error) {
	contents, err := os.ReadFile("C:\\dev\\Wildberries\\OrderService\\configs\\appconfig.yml")
	if err != nil {
		return nil, err
	}
	var config AppConfig
	err = yaml.Unmarshal(contents, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
