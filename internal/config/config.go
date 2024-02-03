package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

// This package parse value from config to struct for now it is only .yml

type Config struct {
	Mongo `yaml:"mongo"`
	HTTP  `yaml:"http"`
}

type Mongo struct {
	Database   string `yaml:"database"`
	Collection `yaml:"collection"`
	URL        string `yaml:"url"`
}
type Collection struct {
	User string `yaml:"user"`
}

type HTTP struct {
	Port string `yaml:"port"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig("config.yml", &cfg); err != nil {
		return nil, fmt.Errorf("error with reading config files %v", err)
	}
	return &cfg, nil
}
