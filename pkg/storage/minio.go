package database

import (
	"net"

	"github.com/go-microservice-boilerplate/pkg/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinio(c *config.CommonConfig) (*minio.Client, error) {
	return minio.New(net.JoinHostPort(c.Minio.Host, c.Minio.Port), &minio.Options{
		Creds:  credentials.NewStaticV4(c.Minio.AccessKeyID, c.Minio.Secret, ""),
		Secure: c.Minio.Secure,
	})
}
