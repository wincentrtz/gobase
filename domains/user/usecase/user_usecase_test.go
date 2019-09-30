package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wincentrtz/gobase/domains/user/mocks"
	"github.com/wincentrtz/gobase/domains/user/usecase"
	"github.com/wincentrtz/gobase/models/responses"
)

const FETCH_USER_BY_ID = "FetchUserById"
const ID = 1
const NAME = "name"
const EMAIL = "email@example.com"

var user = &responses.User{
	ID:    ID,
	Name:  NAME,
	Email: EMAIL,
}

func TestFetchUserById(t *testing.T) {
	mockRepo := new(mocks.Repository)
	mockRepo.On(FETCH_USER_BY_ID, ID).Return(user, nil)
	userUsecase := usecase.NewUserUsecase(mockRepo)
	expectedResult, _ := userUsecase.FetchUserById(ID)
	assert.Equal(t, user, expectedResult)
}
