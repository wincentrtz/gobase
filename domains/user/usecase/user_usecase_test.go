package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wincentrtz/gobase/domains/user/mocks"
	"github.com/wincentrtz/gobase/domains/user/usecase"
	"github.com/wincentrtz/gobase/models/dto/responses"
)

const FetchUserById = "FetchUserById"
const Id = 1
const Name = "name"
const Email = "email@example.com"

var user = &responses.User{
	ID:    Id,
	Name:  Name,
	Email: Email,
}

func TestFetchUserById(t *testing.T) {
	mockRepo := new(mocks.Repository)
	mockRepo.On(FetchUserById, Id).Return(user, nil)
	userUsecase := usecase.NewUserUsecase(mockRepo)
	expectedResult, err := userUsecase.FetchUserById(Id)
	assert.Equal(t, user, expectedResult)
	assert.NotNil(t, expectedResult)
	assert.NoError(t, err)
}
