package v1

import (
	"context"

	"github.com/balcieren/go-microservice/pkg/config"
	"github.com/balcieren/go-microservice/pkg/log"
	"github.com/balcieren/go-microservice/pkg/proto"
	"github.com/balcieren/go-microservice/pkg/utils"
	"github.com/gofiber/fiber/v2"
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

func (h *Handler) ListUsers(c *fiber.Ctx) error {
	resp, err := h.service.user.ListUsers(context.Background(), &proto.ListUsersRequest{
		Page:    int32(c.QueryInt("page", 1)),
		PerPage: int32(c.QueryInt("per_page", 10)),
	})
	if err != nil {
		code, msg := utils.ConvertGRPCErrorToHTTP(err)
		return fiber.NewError(code, msg)
	}

	return c.JSON(resp)
}

func (h *Handler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	resp, err := h.service.user.GetUser(context.Background(), &proto.GetUserRequest{
		Id: id,
	})
	if err != nil {
		code, msg := utils.ConvertGRPCErrorToHTTP(err)
		return fiber.NewError(code, msg)
	}

	return c.JSON(resp)
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var payload CreateUserInput
	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	_, err := h.service.user.CreateUser(context.Background(), &proto.CreateUserRequest{
		UserName: payload.Username,
		Email:    payload.Email,
		Age:      uint64(payload.Age),
	})
	if err != nil {
		code, msg := utils.ConvertGRPCErrorToHTTP(err)
		return fiber.NewError(code, msg)
	}

	return c.Status(fiber.StatusCreated).JSON("user has created successfully")
}

func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	var payload UpdateUserInput
	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	_, err := h.service.user.UpdateUser(context.Background(), &proto.UpdateUserRequest{
		Id:       c.Params("id"),
		UserName: payload.Username,
		Email:    payload.Email,
		Age:      uint64(payload.Age),
	})
	if err != nil {
		code, msg := utils.ConvertGRPCErrorToHTTP(err)
		return fiber.NewError(code, msg)
	}

	return c.Status(fiber.StatusAccepted).JSON("user has updated successfully")
}

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	_, err := h.service.user.DeleteUser(context.Background(), &proto.DeleteUserRequest{
		Id: c.Params("id"),
	})
	if err != nil {
		code, msg := utils.ConvertGRPCErrorToHTTP(err)
		return fiber.NewError(code, msg)
	}

	return c.Status(fiber.StatusNoContent).JSON("user has deleted successfully")
}
