package v1

import (
	"log"

	"github.com/balcieren/go-microservice/pkg/config"
)

type Handler struct {
	service *Service
	config  *config.Config
	log     *log.Logger
}

func NewHandler(s *Service, c *config.Config, l *log.Logger) *Handler {
	return &Handler{
		service: s,
		config:  c,
		log:     l,
	}
}
