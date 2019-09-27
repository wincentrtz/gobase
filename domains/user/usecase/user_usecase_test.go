package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wincentrtz/gobase/domains/user/mocks"
	"github.com/wincentrtz/gobase/domains/user/usecase"
	"github.com/wincentrtz/gobase/models/responses"
)

var user = &responses.User{
	ID:    1,
	Name:  "name",
	Email: "email",
}

func TestCalc(t *testing.T) {
	mockRepo := new(mocks.Repository)
	mockRepo.On("FetchUserById", 1).Return(user, nil)
	uc := usecase.NewUserUsecase(mockRepo)
	res, _ := uc.FetchUserById(1)
	assert.Equal(t, user, res)
}
