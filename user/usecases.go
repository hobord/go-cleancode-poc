package user

import (
	"context"

	"github.com/hobord/go-cleancode-poc/models"
)

// Usecases represent the user's usecases contract
type Usecases interface {
	GetByID(ctx context.Context, id string) (*models.User, error)
	List(ctx context.Context, listOptions models.ListOptions) ([]*models.User, error)
	Create(ctx context.Context, user models.User) (*models.User, error)
	Update(ctx context.Context, user models.User) (*models.User, error)
	Delete(ctx context.Context, id string) error
}

type userUsecases struct {
	repository Repository
}

// NewUserUsecases will create an implementation of user.Usecases
func NewUserUsecases(repository Repository) Usecases {
	return &userUsecases{
		repository: repository,
	}
}

func (s *userUsecases) GetByID(ctx context.Context, id string) (*models.User, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *userUsecases) List(ctx context.Context, listOptions models.ListOptions) ([]*models.User, error) {
	return s.repository.List(ctx, listOptions)
}

func (s *userUsecases) Create(ctx context.Context, user models.User) (*models.User, error) {
	return s.repository.Create(ctx, user)
}

func (s *userUsecases) Update(ctx context.Context, user models.User) (*models.User, error) {
	return s.repository.Update(ctx, user)
}

func (s *userUsecases) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
