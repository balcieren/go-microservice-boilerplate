package v1

import (
	"context"
	"math"

	"entgo.io/ent/dialect/sql"

	"github.com/go-microservice-boilerplate/pkg/ent"
	"github.com/go-microservice-boilerplate/pkg/ent/predicate"
	"github.com/go-microservice-boilerplate/pkg/ent/user"
	"github.com/go-microservice-boilerplate/pkg/log"
	"github.com/go-microservice-boilerplate/pkg/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	proto.UnimplementedUserServiceServer
	db  *ent.Client
	log *log.Logger
}

func New(db *ent.Client, l *log.Logger) *Service {
	return &Service{
		db:  db,
		log: l,
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
		Email:     u.Email,
		Age:       u.Age,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}, nil
}

func (s *Service) CreateUser(ctx context.Context, in *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {

	_, err := s.db.User.Create().
		SetUserName(in.GetUserName()).
		SetEmail(in.GetEmail()).
		SetNillableAge(in.Age).
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

	var username *string = nil
	if in.GetUserName() != "" {
		username = &in.UserName
	}

	var email *string = nil
	if in.GetEmail() != "" {
		email = &in.Email
	}

	query := s.db.User.UpdateOneID(id).
		SetNillableUserName(username).
		SetNillableEmail(email).
		SetNillableAge(in.Age)
	_, err = query.Save(ctx)
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
