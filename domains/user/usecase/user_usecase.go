package usecase

import (
	"github.com/wincentrtz/gobase/domains/user"
	"github.com/wincentrtz/gobase/domains/user/repository"
	"github.com/wincentrtz/gobase/models/responses"
)

type userUsecase struct {
	userRepo user.Repository
}

func NewUserUsecase() user.Usecase {
	return &userUsecase{
		userRepo: repository.NewUserRepository(),
	}
}

func (pu *userUsecase) FetchUserById(userId int) (*responses.User, error) {
	user, err := pu.userRepo.FetchUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
