package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	API_PORT, GRPC_PORT                     string
	PROXY_HOST_NAME, PROXY_PORT             string
	USER_API_HOST_NAME, USER_GRPC_HOST_NAME string
	USER_POSTGRESQL_URI                     string
}

func New() (*Config, error) {
	godotenv.Load("../../.env")

	return &Config{
		API_PORT:            os.Getenv("API_PORT"),
		GRPC_PORT:           os.Getenv("GRPC_PORT"),
		PROXY_HOST_NAME:     os.Getenv("PROXY_HOST_NAME"),
		PROXY_PORT:          os.Getenv("PROXY_PORT"),
		USER_API_HOST_NAME:  os.Getenv("USER_API_HOST_NAME"),
		USER_GRPC_HOST_NAME: os.Getenv("USER_GRPC_HOST_NAME"),
		USER_POSTGRESQL_URI: os.Getenv("USER_POSTGRESQL_URI"),
	}, nil
}
