package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	AppEnv   string `json:"app_env" yaml:"app_env"`
	Postgres struct {
		Host     string `json:"host" yaml:"host"`
		Port     string `json:"port" yaml:"port"`
		User     string `json:"user" yaml:"user"`
		Database string `json:"database" yaml:"database"`
		Password string `json:"password" yaml:"password"`
	} `json:"postgres" yaml:"postgres"`
	Swagger struct {
		Host     string `json:"host" yaml:"host"`
		BasePath string `json:"base_path" yaml:"base_path"`
		User     string `json:"user" yaml:"user"`
		Password string `json:"password" yaml:"password"`
	} `json:"swagger" yaml:"swagger"`
}

func NewEnv() (*Env, error) {
	godotenv.Load()
	env := Env{}

	env.AppEnv = os.Getenv("APP_ENV")
	env.Postgres.Host = os.Getenv("POSTGRES_HOST")
	env.Postgres.Port = os.Getenv("POSTGRES_PORT")
	env.Postgres.User = os.Getenv("POSTGRES_USER")
	env.Postgres.Database = os.Getenv("POSTGRES_DB")
	env.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")

	env.Swagger.Host = os.Getenv("SWAGGER_HOST")
	env.Swagger.BasePath = os.Getenv("SWAGGER_BASE_PATH")
	env.Swagger.User = os.Getenv("SWAGGER_USER")
	env.Swagger.Password = os.Getenv("SWAGGER_PASSWORD")

	return &env, nil
}
