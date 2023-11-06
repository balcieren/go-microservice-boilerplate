package v1

import (
	"context"

	"github.com/go-microservice-boilerplate/pkg/config"
	_ "github.com/go-microservice-boilerplate/pkg/helper"
	"github.com/go-microservice-boilerplate/pkg/log"
	"github.com/go-microservice-boilerplate/pkg/proto"
	"github.com/go-microservice-boilerplate/pkg/utils"
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

// @Summary List Users
// @Description List Users
// @Tags User
// @Accept  json
// @Produce  json
// @Param page query int false "Page"
// @Param per_page query int false "Per Page"
// @Success 200 {object} proto.ListUsersResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Security BearerAuth
// @Router /v1/users [get]
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

// @Summary Get User
// @Description Get User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 200 {object} proto.GetUserResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Security BearerAuth
// @Router /v1/users/{id} [get]
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

// @Summary Create User
// @Description Create User
// @Tags User
// @Accept  json
// @Produce  json
// @Param payload body CreateUserInput true "Create User Body"
// @Success 201 {string} string "user has created successfully"
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Security BearerAuth
// @Router /v1/users [post]
func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var payload CreateUserInput
	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	_, err := h.service.user.CreateUser(context.Background(), &proto.CreateUserRequest{
		UserName: payload.Username,
		Email:    payload.Email,
		Age:      payload.Age,
	})
	if err != nil {
		code, msg := utils.ConvertGRPCErrorToHTTP(err)
		return fiber.NewError(code, msg)
	}

	return c.Status(fiber.StatusCreated).JSON("user has created successfully")
}

// @Summary Update User
// @Description Update User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Param payload body UpdateUserInput true "Update User Body"
// @Success 202 {string} string "user has updated successfully"
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Security BearerAuth
// @Router /v1/users/{id} [patch]
func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	var payload UpdateUserInput
	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	_, err := h.service.user.UpdateUser(context.Background(), &proto.UpdateUserRequest{
		Id:       c.Params("id"),
		UserName: payload.Username,
		Email:    payload.Email,
		Age:      payload.Age,
	})
	if err != nil {
		code, msg := utils.ConvertGRPCErrorToHTTP(err)
		return fiber.NewError(code, msg)
	}

	return c.Status(fiber.StatusAccepted).JSON("user has updated successfully")
}

// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 204 {string} string "user has deleted successfully"
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Security BearerAuth
// @Router /v1/users/{id} [delete]
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
