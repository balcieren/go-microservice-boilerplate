package v1

import (
	"context"
	"errors"

	"github.com/balcieren/go-microservice-boilerplate/pkg/entity"
	"github.com/balcieren/go-microservice-boilerplate/pkg/fail"
	"github.com/balcieren/go-microservice-boilerplate/pkg/proto"
	"github.com/balcieren/go-microservice-boilerplate/pkg/query"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

var (
	ErrOwnerNotFound     = errors.New("owner not found")
	ErrOwnerNameRequired = errors.New("owner name is required")
	ErrOwnerListFailed   = errors.New("failed to list owners")
	ErrOwnerCreateFailed = errors.New("failed to create owner")
	ErrOwnerUpdateFailed = errors.New("failed to update owner")
	ErrOwnerDeleteFailed = errors.New("failed to delete owner")
)

type Service struct {
	proto.UnimplementedOwnerServiceServer
	query *query.Query
}

func NewService(q *query.Query) *Service {
	return &Service{
		query: q,
	}
}

func (s *Service) ListOwners(ctx context.Context, req *proto.ListOwnersRequest) (*proto.ListOwnersResponse, error) {
	oq := s.query.Owner

	owners, count, err := oq.WithContext(ctx).FindByPage((int(req.GetPage())-1)*int(req.GetPerPage()), int(req.GetPerPage()))
	if err != nil {
		return nil, fail.New(codes.NotFound, err.Error())
	}

	rows := make([]*proto.Owner, 0)
	for _, owner := range owners {
		resp := proto.Owner{
			Id:        owner.ID.String(),
			Name:      owner.Name,
			CreatedAt: owner.CreatedAt.Unix(),
			UpdatedAt: owner.UpdatedAt.Unix(),
		}

		rows = append(rows, &resp)
	}

	return &proto.ListOwnersResponse{
		Rows:       rows,
		Page:       req.Page,
		PerPage:    req.PerPage,
		Total:      count,
		TotalPages: (count / req.PerPage) + 1,
	}, nil
}
func (s *Service) GetOwner(ctx context.Context, req *proto.GetOwnerRequest) (*proto.GetOwnerResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, fail.New(codes.InvalidArgument, err.Error())
	}

	oq := s.query.Owner

	owner, err := oq.WithContext(ctx).Where(oq.ID.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fail.New(codes.NotFound, err.Error())
		}
		return nil, fail.New(codes.Internal, err.Error())
	}

	return &proto.GetOwnerResponse{
		Id:        owner.ID.String(),
		Name:      owner.Name,
		CreatedAt: owner.CreatedAt.Unix(),
		UpdatedAt: owner.UpdatedAt.Unix(),
	}, nil
}
func (s *Service) CreateOwner(ctx context.Context, req *proto.CreateOwnerRequest) (*proto.CreateOwnerResponse, error) {
	if len(req.GetName()) == 0 {
		return nil, fail.New(codes.InvalidArgument, "name is required")
	}

	owner := entity.Owner{
		Name: req.Name,
	}

	oq := s.query.Owner

	if err := oq.WithContext(ctx).Create(&owner); err != nil {
		return nil, fail.New(codes.Internal, err.Error())
	}

	return &proto.CreateOwnerResponse{
		Message: "owner created",
	}, nil
}
func (s *Service) UpdateOwner(ctx context.Context, req *proto.UpdateOwnerRequest) (*proto.UpdateOwnerResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, fail.New(codes.InvalidArgument, err.Error())
	}

	fields := make([]field.AssignExpr, 0)

	if req.Name != nil {
		fields = append(fields, s.query.Owner.Name.Value(req.Name.GetValue()))
	}

	oq := s.query.Owner

	_, err = oq.WithContext(ctx).Where(oq.ID.Eq(id)).UpdateSimple(fields...)
	if err != nil {
		return nil, fail.New(codes.Internal, err.Error())
	}

	return &proto.UpdateOwnerResponse{
		Message: "owner updated",
	}, nil
}
func (s *Service) DeleteOwner(ctx context.Context, req *proto.DeleteOwnerRequest) (*proto.DeleteOwnerResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, fail.New(codes.NotFound, err.Error())
	}

	if err := s.query.Transaction(func(tx *query.Query) error {
		oq := tx.Owner
		pq := tx.Pet

		_, err = oq.WithContext(ctx).Where(oq.ID.Eq(id)).Delete()
		if err != nil {
			return fail.New(codes.Internal, err.Error())
		}

		_, err = pq.WithContext(ctx).Where(pq.OwnerID.Eq(id)).UpdateSimple(pq.OwnerID.Value(nil))
		if err != nil {
			return fail.New(codes.Internal, err.Error())
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &proto.DeleteOwnerResponse{
		Message: "owner deleted",
	}, nil
}
