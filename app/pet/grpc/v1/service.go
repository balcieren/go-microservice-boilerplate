package v1

import (
	"context"
	"errors"

	"github.com/balcieren/go-microservice-boilerplate/pkg/entity"
	"github.com/balcieren/go-microservice-boilerplate/pkg/fail"
	"github.com/balcieren/go-microservice-boilerplate/pkg/proto"
	"github.com/balcieren/go-microservice-boilerplate/pkg/query"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

var (
	ErrPetNotFound     = errors.New("pet not found")
	ErrPetTypeInvalid  = errors.New("invalid pet type")
	ErrPetNameRequired = errors.New("pet name is required")
	ErrPetListFailed   = errors.New("failed to list pets")
	ErrPetCreateFailed = errors.New("failed to create pet")
	ErrPetUpdateFailed = errors.New("failed to update pet")
	ErrPetDeleteFailed = errors.New("failed to delete pet")
)

type Service struct {
	proto.UnimplementedPetServiceServer
	query *query.Query
}

func NewService(q *query.Query) *Service {
	return &Service{
		query: q,
	}
}

func (s *Service) ListPets(ctx context.Context, req *proto.ListPetsRequest) (*proto.ListPetsResponse, error) {
	pq := s.query.Pet
	query := pq.WithContext(ctx)

	if req.HasOwner != nil {
		if req.HasOwner.GetValue() {
			query = query.Where(pq.OwnerID.IsNotNull())
		} else {
			query = query.Where(pq.OwnerID.IsNull())
		}
	}

	if req.OwnerId != nil {
		ownerID, err := uuid.Parse(req.OwnerId.GetValue())
		if err != nil {
			return nil, fail.New(fiber.StatusBadRequest, err.Error())
		}

		query = query.Where(pq.OwnerID.Eq(ownerID))
	}

	pets, count, err := query.FindByPage(
		(int(req.GetPage())-1)*int(req.GetPerPage()),
		int(req.GetPerPage()),
	)
	if err != nil {
		return nil, fail.New(fiber.StatusNotFound, ErrPetListFailed.Error())
	}

	rows := make([]*proto.Pet, 0)
	if len(pets) > 0 {
		for _, pet := range pets {
			row := proto.Pet{
				Id:        pet.ID.String(),
				OwnerId:   req.OwnerId,
				Name:      pet.Name,
				Type:      pet.Type.String(),
				CreatedAt: pet.CreatedAt.Unix(),
				UpdatedAt: pet.UpdatedAt.Unix(),
			}

			if pet.OwnerID != nil {
				row.OwnerId = wrapperspb.String(pet.OwnerID.String())
			}

			rows = append(rows, &row)
		}
	}

	return &proto.ListPetsResponse{
		Rows:       rows,
		Page:       req.GetPage(),
		PerPage:    req.GetPerPage(),
		Total:      count,
		TotalPages: (count / req.GetPerPage()) + 1,
	}, nil

}
func (s *Service) GetPet(ctx context.Context, req *proto.GetPetRequest) (*proto.GetPetResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, fail.New(fiber.StatusBadRequest, err.Error())
	}

	pq := s.query.Pet

	pet, err := pq.WithContext(ctx).Where(pq.ID.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fail.New(fiber.StatusNotFound, ErrPetNotFound.Error())
		}
		return nil, err
	}

	resp := proto.GetPetResponse{
		Id:   pet.ID.String(),
		Name: pet.Name,
		Type: pet.Type.String(),
	}

	if pet.OwnerID != nil {
		resp.OwnerId = wrapperspb.String(pet.OwnerID.String())
	}

	return &resp, nil
}
func (s *Service) CreatePet(ctx context.Context, req *proto.CreatePetRequest) (*proto.CreatePetResponse, error) {
	pet := entity.Pet{}

	if len(req.Name) == 0 {
		return nil, fail.New(fiber.StatusBadRequest, ErrPetNameRequired.Error())
	}
	pet.Name = req.Name

	if !entity.PetType(req.Type).IsValid() {
		return nil, fail.New(fiber.StatusBadRequest, ErrPetTypeInvalid.Error())
	}
	pet.Type = entity.PetType(req.Type)

	if req.OwnerId != nil {
		ownerID, err := uuid.Parse(req.GetOwnerId().GetValue())
		if err != nil {
			return nil, fail.New(fiber.StatusBadRequest, err.Error())
		}
		pet.OwnerID = &ownerID
	}

	pq := s.query.Pet

	if err := pq.WithContext(ctx).Create(&pet); err != nil {
		return nil, fail.New(fiber.StatusInternalServerError, ErrPetCreateFailed.Error())
	}

	return &proto.CreatePetResponse{
		Message: "pet created",
	}, nil
}
func (s *Service) UpdatePet(ctx context.Context, req *proto.UpdatePetRequest) (*proto.UpdatePetResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, fail.New(fiber.StatusBadRequest, err.Error())
	}

	fields := make([]field.AssignExpr, 0)

	if req.Name != nil {
		if len(req.GetName().GetValue()) == 0 {
			return nil, fail.New(fiber.StatusBadRequest, ErrPetNameRequired.Error())
		}

		fields = append(fields, s.query.Pet.Name.Value(req.GetName().GetValue()))
	}

	if req.Type != nil {
		if !entity.PetType(req.GetType().GetValue()).IsValid() {
			return nil, fail.New(fiber.StatusBadRequest, ErrPetTypeInvalid.Error())
		}

		fields = append(fields, s.query.Pet.Type.Value(req.GetType().GetValue()))
	}

	if req.OwnerId != nil {
		ownerID, err := uuid.Parse(req.GetOwnerId().GetValue())
		if err != nil {
			return nil, fail.New(fiber.StatusBadRequest, err.Error())
		}

		fields = append(fields, s.query.Pet.OwnerID.Value(ownerID))
	}

	pq := s.query.Pet

	_, err = pq.WithContext(ctx).Where(pq.ID.Eq(id)).UpdateSimple(fields...)
	if err != nil {
		return nil, fail.New(fiber.StatusInternalServerError, ErrPetUpdateFailed.Error())
	}

	return &proto.UpdatePetResponse{
		Message: "pet updated",
	}, nil
}
func (s *Service) DeletePet(ctx context.Context, req *proto.DeletePetRequest) (*proto.DeletePetResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, fail.New(fiber.StatusBadRequest, err.Error())
	}

	pq := s.query.Pet

	_, err = pq.WithContext(ctx).Where(pq.ID.Eq(id)).Delete()
	if err != nil {
		return nil, fail.New(fiber.StatusInternalServerError, ErrPetDeleteFailed.Error())
	}

	return &proto.DeletePetResponse{
		Message: "pet deleted",
	}, nil
}
