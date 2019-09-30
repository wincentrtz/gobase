package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/wincentrtz/gobase/models/responses"
)

type Usecase struct {
	mock.Mock
}

func (_m *Usecase) FetchUserById(userId int) (responses.User, error) {
	ret := _m.Called(userId)
	return ret.Get(0).(responses.User), ret.Error(1)
}
