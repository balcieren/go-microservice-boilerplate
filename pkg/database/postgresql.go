package database

import (
	"fmt"

	"github.com/balcieren/go-microservice-boilerplate/pkg/config"
	"github.com/balcieren/go-microservice-boilerplate/pkg/entity"
	"github.com/balcieren/go-microservice-boilerplate/pkg/query"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgreSQL(env *config.Env) (*gorm.DB, *query.Query, error) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		env.Postgres.Host,
		env.Postgres.Port,
		env.Postgres.User,
		env.Postgres.Database,
		env.Postgres.Password,
	)), &gorm.Config{})
	if err != nil {
		return nil, nil, fmt.Errorf("failed opening connection to postgres: %v", err.Error())
	}

	db.AutoMigrate(
		&entity.Owner{},
		&entity.Pet{},
	)

	return db, query.Use(db), nil
}
