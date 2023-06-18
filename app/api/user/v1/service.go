package v1

import (
	"net"

	"github.com/balcieren/go-microservice/pkg/config"
	"github.com/balcieren/go-microservice/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	user proto.UserServiceClient
}

func NewService(c *config.Config) (*Service, error) {
	conn, err := grpc.Dial(
		net.JoinHostPort(c.USER_GRPC_HOST_NAME, c.GRPC_PORT),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Service{
		user: proto.NewUserServiceClient(conn),
	}, nil
}
