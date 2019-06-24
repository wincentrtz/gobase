package usecase

import (
	"database/sql"

	userRepository "github.com/wincentrtz/gobase/domains/user/repository"

	"github.com/wincentrtz/gobase/domains/user"
	"github.com/wincentrtz/gobase/models"
)

type userUsecase struct {
	userRepo user.Repository
}

func NewUserUsecase(db *sql.DB) user.Usecase {
	return &userUsecase{
		userRepo: userRepository.NewUserRepository(db),
	}
}

func (pu *userUsecase) FetchUserById(userId int) (*models.User, error) {
	user, err := pu.userRepo.FetchUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
