package usecase

import (
	"github.com/wincentrtz/gobase/domains/user"
	"github.com/wincentrtz/gobase/models/dto/responses"
)

type userUsecase struct {
	userRepo user.Repository
}

func NewUserUsecase(ur user.Repository) user.Usecase {
	return &userUsecase{
		userRepo: ur,
	}
}

func (pu *userUsecase) FetchUserById(userId int) (*responses.User, error) {
	user, err := pu.userRepo.FetchUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
