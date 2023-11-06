package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	API struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"api"`
	GRPC struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"grpc"`
	Proxy struct {
		Port string `json:"port"`
	} `json:"proxy"`
}

func New() (*Config, error) {
	godotenv.Load("../../.env")

	return &Config{
		API: struct {
			Host string `json:"host"`
			Port string `json:"port"`
		}{
			Host: os.Getenv("API_HOST_NAME"),
			Port: os.Getenv("API_PORT"),
		},

		GRPC: struct {
			Host string `json:"host"`
			Port string `json:"port"`
		}{
			Host: os.Getenv("GRPC_HOST_NAME"),
			Port: os.Getenv("GRPC_PORT"),
		},

		Proxy: struct {
			Port string `json:"port"`
		}{
			Port: os.Getenv("PROXY_PORT"),
		},
	}, nil
}
