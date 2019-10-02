package template

import (
	"github.com/stretchr/testify/mock"
)

type Usecase struct {
	mock.Mock
}

func (_m *Usecase) ExampleUsecaseMocks(userId int) (int, error) {
	ret := _m.Called(userId)
	return 1, ret.Error(1)
}
