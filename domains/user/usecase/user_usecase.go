package usecase

import (
	userRepository "github.com/wincentrtz/gobase/domains/user/repository"

	"github.com/wincentrtz/gobase/domains/user"
	"github.com/wincentrtz/gobase/models"
)

type userUsecase struct {
	userRepo user.Repository
}

func NewUserUsecase() user.Usecase {
	return &userUsecase{
		userRepo: userRepository.NewUserRepository(),
	}
}

func (pu *userUsecase) FetchUserById(userId int) (*models.User, error) {
	user, err := pu.userRepo.FetchUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
