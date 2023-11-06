package worker

import (
	"net"

	"github.com/go-microservice-boilerplate/pkg/config"
	"github.com/hibiken/asynq"
)

func NewAsynq(c *config.CommonConfig) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{
		Addr:     net.JoinHostPort(c.Redis.Port, c.Redis.Port),
		Password: c.Redis.Password,
		DB:       2,
	})
}
