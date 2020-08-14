package user

import "github.com/wincentrtz/gobase/models/dto/responses"

type Usecase interface {
	FetchUserById(userId int) (*responses.UserResponse, error)
}
