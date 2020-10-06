package repository

import (
	"github.com/wincentrtz/gobase/domains/user"
	"github.com/wincentrtz/gobase/models/dto/responses"
	"github.com/wincentrtz/gobase/models/entity"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) FetchUserById(userId int) (*responses.UserResponse, error) {

	user := entity.User{}
	ur.db.First(&user, userId)

	userResponse := &responses.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return userResponse, nil
}
