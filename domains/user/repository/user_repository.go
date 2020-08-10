package repository

import (
	"database/sql"
	"strconv"

	"github.com/wincentrtz/gobase/domains/user"
	"github.com/wincentrtz/gobase/gobase/utils"
	"github.com/wincentrtz/gobase/models/dto/responses"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) user.Repository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) FetchUserById(userId int) (*responses.User, error) {
	var id int
	var name string
	var email string

	query := utils.NewQueryBuilder().
		Table("users").
		Select("id,name,email").
		Where("id", "=", strconv.Itoa(userId)).
		Build()

	err := ur.db.QueryRow(query).Scan(
		&id,
		&name,
		&email,
	)

	if err != nil {
		return nil, err
	}

	user := &responses.User{
		ID:    id,
		Name:  name,
		Email: email,
	}

	return user, nil
}
