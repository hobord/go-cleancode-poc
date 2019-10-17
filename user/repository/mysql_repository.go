package repository

import (
	"context"
	"database/sql"

	uuid "github.com/google/uuid"

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
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	user := &models.User{}

	err = row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *mysqlUserRepo) getMultiple(ctx context.Context, query string, args ...interface{}) ([]*models.User, error) {
	users := make([]*models.User, 1)
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.Name,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (m *mysqlUserRepo) GetByID(ctx context.Context, id string) (*models.User, error) {
	query := `SELECT id, email, password, name, created_at, updated_at FROM user WHERE id=?`
	return m.getOne(ctx, query, id)
}

func (m *mysqlUserRepo) List(ctx context.Context, listOptions models.ListOptions) ([]*models.User, error) {
	query := `SELECT id, email, password, name, created_at, updated_at FROM user`

	query = query + listOptions.MysqlQueryStr()

	return m.getMultiple(ctx, query)
}

func (m *mysqlUserRepo) Create(ctx context.Context, user models.User) (*models.User, error) {
	user.ID = uuid.New().String()
	return nil, nil
}

func (m *mysqlUserRepo) Update(ctx context.Context, user models.User) (*models.User, error) {
	return nil, nil
}

func (m *mysqlUserRepo) Delete(ctx context.Context, id string) error {
	return nil
}
