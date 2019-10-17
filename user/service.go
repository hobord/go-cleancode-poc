package user

import (
	"context"

	"github.com/hobord/go-cleancode-poc/models"
)

// Service represent the user's service contract
type Service interface {
	GetByID(ctx context.Context, id string) (*models.User, error)
	List(ctx context.Context, listOptions models.ListOptions) ([]*models.User, error)
	Create(ctx context.Context, user models.User) (*models.User, error)
	Update(ctx context.Context, user models.User) (*models.User, error)
	Delete(ctx context.Context, id string) error
}

type userService struct {
	Repository *Repository
}

// NewUserService will create an implementation of user.Service
func NewUserService(repository *Repository) *Service {
	return &userService{
		Repository: repository,
	}
}

func (s *userService) GetByID(ctx context.Context, id string) (*models.User, error) {
	return nil, nil
}

func (s *userService) List(ctx context.Context, listOptions models.ListOptions) ([]*models.User, error) {
	return nil, nil
}

func (s *userService) Create(ctx context.Context, user models.User) (*models.User, error) {
	return nil, nil
}

func (s *userService) Update(ctx context.Context, user models.User) (*models.User, error) {
	return nil, nil
}

func (s *userService) Delete(ctx context.Context, id string) error {
	return nil
}
