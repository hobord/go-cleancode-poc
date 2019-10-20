package user

import (
	"context"

	"github.com/hobord/go-cleancode-poc/models"
)

// Service represent the user's (usecases) service contract
type Service interface {
	GetByID(ctx context.Context, id string) (*models.User, error)
	List(ctx context.Context, listOptions models.ListOptions) ([]*models.User, error)
	Create(ctx context.Context, user models.User) (*models.User, error)
	Update(ctx context.Context, user models.User) (*models.User, error)
	Delete(ctx context.Context, id string) error
}

type userService struct {
	repository Repository
}

// NewUserService will create an implementation of user.Service
func NewUserService(repository Repository) Service {
	return &userService{
		repository: repository,
	}
}

func (s *userService) GetByID(ctx context.Context, id string) (*models.User, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *userService) List(ctx context.Context, listOptions models.ListOptions) ([]*models.User, error) {
	return s.repository.List(ctx, listOptions)
}

func (s *userService) Create(ctx context.Context, user models.User) (*models.User, error) {
	return s.repository.Create(ctx, user)
}

func (s *userService) Update(ctx context.Context, user models.User) (*models.User, error) {
	return s.repository.Update(ctx, user)
}

func (s *userService) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
