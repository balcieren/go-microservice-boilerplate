package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Common struct {
	API struct {
		HOST string `json:"host" yaml:"host"`
		PORT string `json:"port" yaml:"port"`
	} `json:"api" yaml:"api"`
	GRPC struct {
		HOST string `json:"host" yaml:"host"`
		PORT string `json:"port" yaml:"port"`
	} `json:"grpc" yaml:"grpc"`
}

func NewCommon() (*Common, error) {
	common := Common{}

	filename, err := filepath.Abs("./config.yaml")
	if err != nil {
		return nil, err
	}

	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(yamlFile, &common); err != nil {
		return nil, err
	}

	return &common, nil
}
