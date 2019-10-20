package repository_test

import (
	"context"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/hobord/go-cleancode-poc/user/repository"
)

func TestMysqlGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "email", "name", "updated_at", "created_at"}).
		AddRow("TEST_ID", "test@test.com", "Mickie Mouse", time.Now(), time.Now())

	query := "SELECT id, email, name, created_at, updated_at FROM user WHERE id=\\?"

	prep := mock.ExpectPrepare(query)
	userID := "TEST_ID"
	prep.ExpectQuery().WithArgs(userID).WillReturnRows(rows)

	a := repository.NewMysqlUserRepository(db)

	anUser, err := a.GetByID(context.TODO(), userID)
	assert.NoError(t, err)
	assert.NotNil(t, anUser)
}
