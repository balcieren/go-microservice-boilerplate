package v1

import (
	"github.com/balcieren/go-microservice-boilerplate/pkg/config"
	"github.com/balcieren/go-microservice-boilerplate/pkg/entity"
	"github.com/balcieren/go-microservice-boilerplate/pkg/fail"
	"github.com/balcieren/go-microservice-boilerplate/pkg/proto"
	utils "github.com/balcieren/go-microservice-boilerplate/pkg/util"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type API struct {
	service proto.PetServiceClient
	router  fiber.Router
}

func NewAPI(env *config.Env, c *grpc.ClientConn, app *fiber.App) *API {
	var path string = "/v1/pets"
	if env.AppEnv == "dev" {
		path = "/api/v1/pets"
	}

	return &API{
		service: proto.NewPetServiceClient(c),
		router:  app.Group(path),
	}
}

func (api API) Setup() {
	api.router.Get("/", api.ListPets)
	api.router.Get("/:id", api.GetPet)
	api.router.Post("/", api.CreatePet)
	api.router.Patch("/:id", api.UpdatePet)
	api.router.Delete("/:id", api.DeletePet)
}

// @ID ListPets
// @Summary List Pets
// @Description List Pets
// @Tags pets v1
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param per_page query int false "Per Page"
// @Param has_owner query bool false "Has Owner"
// @Param owner_id query string false "Owner ID"
// @Success 200 {object} ListPetsResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /v1/pets [get]
func (api API) ListPets(c *fiber.Ctx) error {
	req := proto.ListPetsRequest{
		Page:    int64(c.QueryInt("page", 1)),
		PerPage: int64(c.QueryInt("per_page", 10)),
	}

	if c.Query("has_owner", "") != "" {
		req.HasOwner = wrapperspb.Bool(c.Query("has_owner") == "true")
	}

	if c.Query("owner_id", "") != "" {
		req.OwnerId = wrapperspb.String(c.Query("owner_id"))
	}

	resp, err := api.service.ListPets(c.Context(), &req)
	if err != nil {
		return fiber.NewError(fail.Convert(err))
	}

	rows := make([]GetPetResponse, 0)
	for _, row := range resp.GetRows() {
		rows = append(rows, GetPetResponse{
			ID:        row.GetId(),
			Name:      row.GetName(),
			Type:      entity.PetType(row.GetType()),
			OwnerID:   utils.Ptr(row.GetOwnerId().GetValue()),
			CreatedAt: row.GetCreatedAt(),
			UpdatedAt: row.GetUpdatedAt(),
		})
	}

	return c.JSON(ListPetsResponse{
		Rows:       rows,
		Page:       int(resp.GetPage()),
		PerPage:    int(resp.GetPerPage()),
		Total:      int(resp.GetTotal()),
		TotalPages: int(resp.GetTotalPages()),
	})
}

// @ID GetPet
// @Summary Get a pet
// @Description Get a pet
// @Tags pets v1
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} GetPetResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 404 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /v1/pets/{id} [get]
func (api API) GetPet(c *fiber.Ctx) error {
	id := c.Params("id")

	resp, err := api.service.GetPet(c.Context(), &proto.GetPetRequest{
		Id: id,
	})
	if err != nil {
		return fiber.NewError(fail.Convert(err))
	}

	return c.JSON(GetPetResponse{
		ID:        resp.GetId(),
		Name:      resp.GetName(),
		Type:      entity.PetType(resp.GetType()),
		CreatedAt: resp.GetCreatedAt(),
		UpdatedAt: resp.GetUpdatedAt(),
		OwnerID:   utils.Ptr(resp.GetOwnerId().GetValue()),
	})
}

// @ID CreatePet
// @Summary Create a pet
// @Description Create a pet
// @Tags pets v1
// @Accept json
// @Produce json
// @Param body body CreatePetBody true "Name"
// @Success 200 {object} CreatePetResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /v1/pets [post]
func (api API) CreatePet(c *fiber.Ctx) error {
	var body CreatePetBody
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	req := proto.CreatePetRequest{
		Name: body.Name,
		Type: body.Type,
	}

	if body.OwnerID != nil {
		req.OwnerId = wrapperspb.String(*body.OwnerID)
	}

	resp, err := api.service.CreatePet(c.Context(), &req)
	if err != nil {
		return fiber.NewError(fail.Convert(err))
	}

	return c.Status(fiber.StatusCreated).JSON(CreatePetResponse{
		Message: resp.GetMessage(),
	})
}

// @ID UpdatePet
// @Summary Update a pet
// @Description Update a pet
// @Tags pets v1
// @Accept json
// @Produce json
// @Param body body UpdatePetBody true "Body"
// @Param owner_id body string false "Owner ID"
// @Success 200 {object} UpdatePetResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /v1/pets/{id} [patch]
func (api API) UpdatePet(c *fiber.Ctx) error {
	id := c.Params("id")

	body := UpdatePetBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	req := proto.UpdatePetRequest{
		Id: id,
	}

	if body.Name != nil {
		req.Name = wrapperspb.String(*body.Name)
	}

	if body.Type != nil {
		req.Type = wrapperspb.String(*body.Type)
	}

	if body.OwnerID != nil {
		req.OwnerId = wrapperspb.String(*body.OwnerID)
	}

	resp, err := api.service.UpdatePet(c.Context(), &req)
	if err != nil {
		return fiber.NewError(fail.Convert(err))
	}

	return c.Status(fiber.StatusAccepted).JSON(UpdatePetResponse{
		Message: resp.GetMessage(),
	})
}

// @ID DeletePet
// @Summary Delete a pet
// @Description Delete a pet
// @Tags pets v1
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} DeletePetResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /v1/pets/{id} [delete]
func (api API) DeletePet(c *fiber.Ctx) error {
	id := c.Params("id")

	resp, err := api.service.DeletePet(c.Context(), &proto.DeletePetRequest{
		Id: id,
	})
	if err != nil {
		return fiber.NewError(fail.Convert(err))
	}

	return c.Status(fiber.StatusNoContent).JSON(DeletePetResponse{
		Message: resp.GetMessage(),
	})
}
