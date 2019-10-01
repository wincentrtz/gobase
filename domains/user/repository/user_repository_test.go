package repository_test

import (
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/wincentrtz/gobase/domains/user/repository"
	"github.com/wincentrtz/gobase/gobase/utils"
	"github.com/wincentrtz/gobase/models/responses"
)

const (
	ID    = 1
	NAME  = "name"
	EMAIL = "email@example.com"
)

var ROWS = []string{"id", "name", "email"}

var mockUser = &responses.User{
	ID:    ID,
	Name:  NAME,
	Email: EMAIL,
}

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	rows := sqlmock.NewRows(ROWS).AddRow(mockUser.ID, mockUser.Name, mockUser.Email)
	query := utils.NewQueryBuilder().
		Table("users").
		Select("id,name,email").
		Where("id", "=", strconv.Itoa(ID)).
		Build()

	mock.ExpectQuery(query).WillReturnRows(rows)
	ur := repository.NewUserRepository(db)

	res, err := ur.FetchUserById(ID)
	assert.Equal(t, res, mockUser)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
