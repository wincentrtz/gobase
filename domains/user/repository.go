package user

import "github.com/wincentrtz/gobase/models/dto/responses"

type Repository interface {
	FetchUserById(userId int) (*responses.User, error)
}
