package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		User   string `yaml:"user"`
		Pass   string `yaml:"pass"`
		DBName string `yaml:"dbname"`
		Host   string `yaml:"host"`
		Port   string `yaml:"port"`
	} `yaml:"database"`
}

func ReadConfig() *Config {
	f, err := os.Open("config.yaml")
	if err != nil {
		log.Fatal("Can't read your config.yaml, err=", err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal("Can't read your config.yaml, err=", err)
	}
	return &cfg
}
