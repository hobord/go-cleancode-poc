package user

import (
	"context"

	"github.com/hobord/go-cleancode-poc/models"
)

// Repository represent the user's repository contract
type Repository interface {
	GetByID(ctx context.Context, id string) (*models.User, error)
	List(ctx context.Context, listOptions models.ListOptions) ([]*models.User, error)
	Create(ctx context.Context, user models.User) (*models.User, error)
	Update(ctx context.Context, user models.User) (*models.User, error)
	Delete(ctx context.Context, id string) error
}
