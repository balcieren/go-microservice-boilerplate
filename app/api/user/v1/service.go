package v1

import (
	"github.com/go-microservice-boilerplate/pkg/config"
	"github.com/go-microservice-boilerplate/pkg/proto"
	"google.golang.org/grpc"
)

type Service struct {
	user proto.UserServiceClient
}

func NewService(c *config.Config, conn *grpc.ClientConn) *Service {
	return &Service{
		user: proto.NewUserServiceClient(conn),
	}
}
