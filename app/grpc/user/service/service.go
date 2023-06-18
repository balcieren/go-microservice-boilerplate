package service

import (
	"context"
	"log"
	"math"

	"entgo.io/ent/dialect/sql"
	"github.com/balcieren/go-microservice/app/grpc/user/ent"
	"github.com/balcieren/go-microservice/app/grpc/user/ent/predicate"
	"github.com/balcieren/go-microservice/app/grpc/user/ent/user"
	"github.com/balcieren/go-microservice/pkg/config"
	"github.com/balcieren/go-microservice/pkg/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	proto.UnimplementedUserServiceServer
	db  *ent.Client
	log *log.Logger
	cfg *config.Config
}

func New(db *ent.Client, l *log.Logger, c *config.Config) *Service {
	return &Service{
		db:  db,
		log: l,
		cfg: c,
	}
}

func (s *Service) ListUsers(ctx context.Context, in *proto.ListUsersRequest) (*proto.ListUsersResponse, error) {
	totalCount, err := s.db.User.Query().Count(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	users, err := s.db.User.Query().Limit(int(in.GetPerPage())).Offset(int(in.GetPerPage()) * (int(in.GetPage()) - 1)).All(ctx)
	if ent.IsNotFound(err) {
		return &proto.ListUsersResponse{}, nil
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var rows []*proto.User
	for _, user := range users {
		rows = append(rows, &proto.User{
			Id:        user.ID.String(),
			UserName:  user.UserName,
			Email:     user.Email,
			Age:       user.Age,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		})
	}

	return &proto.ListUsersResponse{
		Rows:       rows,
		Total:      int32(totalCount),
		TotalPages: int32(math.Ceil(float64(totalCount) / float64(in.GetPerPage()))),
	}, nil
}

func (s *Service) GetUser(ctx context.Context, in *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	var id predicate.User = func(s *sql.Selector) {}
	if in.GetId() != "" {
		uid, err := uuid.Parse(in.GetId())
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		id = user.ID(uid)
	}

	u, err := s.db.User.Query().Where(id).First(ctx)
	if ent.IsNotFound(err) {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetUserResponse{
		Id:        u.ID.String(),
		UserName:  u.UserName,
		Age:       u.Age,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}, nil
}

func (s *Service) CreateUser(ctx context.Context, in *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	var age *uint64 = nil
	if in.GetAge() != 0 {
		age = &in.Age
	}

	_, err := s.db.User.Create().
		SetUserName(in.GetUserName()).
		SetEmail(in.GetEmail()).
		SetNillableAge(age).
		Save(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreateUserResponse{}, nil
}

func (s *Service) UpdateUser(ctx context.Context, in *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	id, err := uuid.Parse(in.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	_, err = s.db.User.UpdateOneID(id).
		SetNillableUserName(&in.Email).
		SetNillableEmail(&in.Email).
		SetNillableAge(&in.Age).
		Save(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.UpdateUserResponse{}, nil
}

func (s *Service) DeleteUser(ctx context.Context, in *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	id, err := uuid.Parse(in.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = s.db.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.DeleteUserResponse{}, nil
}
