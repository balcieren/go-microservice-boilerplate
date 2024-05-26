package v1

import (
	"fmt"

	"github.com/balcieren/go-microservice-boilerplate/pkg/fail"
	"github.com/balcieren/go-microservice-boilerplate/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service proto.OwnerServiceClient
}

func NewHandler(c *grpc.ClientConn) *Handler {
	fmt.Println(c.GetState().String())
	return &Handler{
		service: proto.NewOwnerServiceClient(c),
	}
}

// @ID ListOwners
// @Summary List owners
// @Description List owners
// @Tags owners v1
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param per_page query int false "Per Page"
// @Success 200 {object} ListOwnersResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /v1/owners [get]
func (h Handler) ListOwners(c *fiber.Ctx) error {
	resp, err := h.service.ListOwners(c.Context(), &proto.ListOwnersRequest{
		Page:    int64(c.QueryInt("page", 1)),
		PerPage: int64(c.QueryInt("per_page", 10)),
	})
	if err != nil {
		return fiber.NewError(fail.Convert(err))
	}

	rows := make([]GetOwnerResponse, 0)
	for _, row := range resp.GetRows() {
		rows = append(rows, GetOwnerResponse{
			ID:        row.GetId(),
			Name:      row.GetName(),
			CreatedAt: row.GetCreatedAt(),
			UpdatedAt: row.GetUpdatedAt(),
		})
	}

	return c.JSON(ListOwnersResponse{
		Rows:       rows,
		Page:       int(resp.GetPage()),
		PerPage:    int(resp.GetPerPage()),
		Total:      int(resp.GetTotal()),
		TotalPages: int(resp.GetTotalPages()),
	})
}

// @ID GetOwner
// @Summary Get a Owner
// @Description Get a Owner
// @Tags owners v1
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} GetOwnerResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 404 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /v1/owners/{id} [get]
func (h Handler) GetOwner(c *fiber.Ctx) error {
	id := c.Params("id")

	resp, err := h.service.GetOwner(c.Context(), &proto.GetOwnerRequest{
		Id: id,
	})
	if err != nil {
		return fiber.NewError(fail.Convert(err))
	}

	return c.JSON(GetOwnerResponse{
		ID:        resp.GetId(),
		Name:      resp.GetName(),
		CreatedAt: resp.GetCreatedAt(),
		UpdatedAt: resp.GetUpdatedAt(),
	})
}

// @ID CreateOwner
// @Summary Create a owner
// @Description Create a owner
// @Tags owners v1
// @Accept json
// @Produce json
// @Param body body CreateOwnerBody true "Body"
// @Success 200 {object} CreateOwnerResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /v1/owners [post]
func (h Handler) CreateOwner(c *fiber.Ctx) error {
	body := CreateOwnerBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	resp, err := h.service.CreateOwner(c.Context(), &proto.CreateOwnerRequest{
		Name: body.Name,
	})
	if err != nil {
		return fiber.NewError(fail.Convert(err))
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

// @ID UpdateOwner
// @Summary Update a owner
// @Description Update a owner
// @Tags owners v1
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param body body UpdateOwnerBody true "Body"
// @Success 200 {object} UpdateOwnerResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /v1/owners/{id} [patch]
func (h Handler) UpdateOwner(c *fiber.Ctx) error {
	body := UpdateOwnerBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	req := &proto.UpdateOwnerRequest{
		Id: c.Params("id"),
	}

	if body.Name != nil {
		req.Name = wrapperspb.String(*body.Name)
	}

	resp, err := h.service.UpdateOwner(c.Context(), req)
	if err != nil {
		return fiber.NewError(fail.Convert(err))
	}

	return c.Status(fiber.StatusAccepted).JSON(UpdateOwnerResponse{
		Message: resp.GetMessage(),
	})
}

// @ID DeleteOwner
// @Summary Delete a owner
// @Description Delete a owner
// @Tags owners v1
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} DeleteOwnerResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /v1/owners/{id} [delete]
func (h Handler) DeleteOwner(c *fiber.Ctx) error {
	id := c.Params("id")

	resp, err := h.service.DeleteOwner(c.Context(), &proto.DeleteOwnerRequest{
		Id: id,
	})
	if err != nil {
		return fiber.NewError(fail.Convert(err))
	}

	return c.Status(fiber.StatusNoContent).JSON(DeleteOwnerResponse{
		Message: resp.GetMessage(),
	})
}
