package repository

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"

	"github.com/hobord/go-cleancode-poc/models"
	"github.com/hobord/go-cleancode-poc/user"
)

type mysqlUserRepo struct {
	DB *sql.DB
}

// NewMysqlUserRepository will create an implementation of user.Repository
func NewMysqlUserRepository(db *sql.DB) user.Repository {
	return &mysqlUserRepo{
		DB: db,
	}
}

func (m *mysqlUserRepo) getOne(ctx context.Context, query string, args ...interface{}) (*models.User, error) {

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	a := &models.User{}

	err = row.Scan(
		&a.ID,
		&a.Email,
		&a.Password,
		&a.Name,
		&a.CreatedAt,
		&a.UpdatedAt,
	)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return a, nil
}

func (m *mysqlUserRepo) GetByID(ctx context.Context, id string) (*models.User, error) {
	query := `SELECT id, name, created_at, updated_at FROM user WHERE id=?`
	return m.getOne(ctx, query, id)
}

func (m *mysqlUserRepo) List(ctx context.Context, listOptions models.ListOptions) ([]*models.User, error) {
	return nil, nil
}

func (m *mysqlUserRepo) Create(ctx context.Context, user models.User) (*models.User, error) {
	return nil, nil
}

func (m *mysqlUserRepo) Update(ctx context.Context, user models.User) (*models.User, error) {
	return nil, nil
}

func (m *mysqlUserRepo) Delete(ctx context.Context, id string) error {
	return nil
}
