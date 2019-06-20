package repository

import (
	"database/sql"
	"strconv"

	"github.com/wincentrtz/gobase/domains/user"
	"github.com/wincentrtz/gobase/models"
	"github.com/wincentrtz/gobase/querybuilder"
)

type userRepository struct {
	Conn *sql.DB
}

func NewUserRepository(Conn *sql.DB) user.Repository {
	return &userRepository{
		Conn,
	}
}

func (m *userRepository) FetchUserById(userId int) (*models.User, error) {
	var id int
	var name string
	var email string

	query := querybuilder.NewQueryBuilder().
		Table("users").
		Select("id,name,email").
		Where("id", "LIKE", "%"+strconv.Itoa(userId)+"%").
		Excecute()

	err := m.Conn.QueryRow(query).Scan(
		&id,
		&name,
		&email,
	)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:    id,
		Name:  name,
		Email: email,
	}

	return user, nil
}
