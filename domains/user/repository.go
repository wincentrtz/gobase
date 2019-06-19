package user

import "github.com/wincentrtz/gobase/models"

type Repository interface {
	FetchUserById(userId int) (*models.User, error)
}
