package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/wincentrtz/gobase/models/dto/responses"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) FetchUserById(userId int) (*responses.UserResponse, error) {
	ret := _m.Called(userId)
	return ret.Get(0).(*responses.UserResponse), ret.Error(1)
}
