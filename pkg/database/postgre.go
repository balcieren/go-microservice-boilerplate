package database

import (
	"context"
	"fmt"

	"github.com/go-microservice-boilerplate/pkg/config"
	"github.com/go-microservice-boilerplate/pkg/ent"
	_ "github.com/lib/pq"
)

func NewPostgreSQL(c *config.CommonConfig) (*ent.Client, error) {
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.User,
		c.Postgres.Database,
		c.Postgres.Password,
	))

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err.Error())
	}
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %v", err.Error())
	}

	return client, nil
}
