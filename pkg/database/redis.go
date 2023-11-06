package database

import (
	"net"
	"strconv"

	"github.com/go-microservice-boilerplate/pkg/config"
	"github.com/go-redis/redis/v8"
)

func NewRedis(c *config.CommonConfig) *redis.Client {
	db, _ := strconv.Atoi(c.Redis.DB)

	rdb := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(c.Redis.Host, c.Redis.Port),
		Password: c.Redis.Password,
		DB:       db,
	})

	return rdb
}
