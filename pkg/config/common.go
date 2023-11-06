package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type CommonConfig struct {
	Postgres struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"postgres"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		DB       string `yaml:"db"`
	} `yaml:"redis"`
	Meili struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Key  string `yaml:"key"`
	} `yaml:"meili"`
	Minio struct {
		Host        string `yaml:"host"`
		Port        string `yaml:"port"`
		AccessKeyID string `yaml:"access_key_id"`
		Secret      string `yaml:"secret"`
		Secure      bool   `yaml:"secure"`
		StaticHost  string `yaml:"static_host"`
		Bucket      string `yaml:"bucket"`
	} `yaml:"minio"`
}

func NewCommonConfig() *CommonConfig {
	config := CommonConfig{}
	filename, _ := filepath.Abs("./config.yaml")
	yamlFile, _ := os.ReadFile(filename)
	yaml.Unmarshal(yamlFile, &config)

	return &config
}
