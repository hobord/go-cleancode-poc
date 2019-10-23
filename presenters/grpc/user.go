package user_grpc

import (
	context "context"
	"errors"

	"github.com/hobord/go-cleancode-poc/model"
	"github.com/hobord/go-cleancode-poc/user"
)

// GrpcServer for user usecases
type GrpcServer struct {
	userUsecases user.Usecases
}

func (s *GrpcServer) getUseCases() user.Usecases {
	return s.userUsecases
}

// CreateGrpcServer is a constructor for GrpcServer
func CreateGrpcServer(userUsecases user.Usecases) (*GrpcServer, error) {
	if userUsecases == nil {
		return nil, errors.New("We need some user usecases")
	}
	return &GrpcServer{
		userUsecases: userUsecases,
	}, nil
}

func (s *GrpcServer) GetByID(ctx context.Context, userIdMessage *UserIdMessage) (*User, error) {
	u, err := s.userUsecases.GetByID(ctx, userIdMessage.Id)

	if err != nil {
		return nil, err
	}

	createdAt := &Timestamp{
		Seconds: u.CreatedAt.Unix(),
	}

	updatedAt := &Timestamp{
		Seconds: u.UpdatedAt.Unix(),
	}

	return &User{
		Id:        u.ID,
		Email:     u.Email,
		Name:      u.Name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, err
}

func (s *GrpcServer) Create(ctx context.Context, user *User) (*User, error) {
	newUser := &model.User{
		user.Email,
		user.Name,
	}
	u, err := s.userUsecases.Create(ctx, newUser)

	if err != nil {
		return nil, err
	}

	createdAt := &Timestamp{
		Seconds: u.CreatedAt.Unix(),
	}

	updatedAt := &Timestamp{
		Seconds: u.UpdatedAt.Unix(),
	}

	return &User{
		Id:        u.ID,
		Email:     u.Email,
		Name:      u.Name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, err
}
