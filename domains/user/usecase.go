package user

import "github.com/wincentrtz/gobase/models"

type Usecase interface {
	FetchUserById(userId int) (*models.User, error)
}
