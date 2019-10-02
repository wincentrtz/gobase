package template

import (
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) ExampleRepositoryMocks(userId int) (int, error) {
	ret := _m.Called(userId)
	return 1, ret.Error(1)
}
