package database

import (
	"context"
	"fmt"

	"github.com/balcieren/go-microservice/app/grpc/user/ent"
	_ "github.com/lib/pq"
)

func NewPostgreSQL(uri string) (*ent.Client, error) {
	client, err := ent.Open("postgres", uri)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %v", err.Error())
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err.Error())
	}

	return client, nil
}
