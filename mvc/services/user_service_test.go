package services

import (
	"github.com/stretchr/testify/assert"
	"microservices/mvc/domain"
	"microservices/mvc/utils"
	"net/http"
	"testing"
)

var (
	userDaoMock usersDaoMock
	getUserFunction = func(userID int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct {

}

func (m* usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestUserServiceNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    "user 0 does not exists",
			StatusCode: http.StatusNotFound,
			Code:       "",
		}
	}
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 does not exists", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userID int64) (user *domain.User, applicationError *utils.ApplicationError) {
		return &domain.User{
			Id:        123,
			FirstName: "Frank",
			LastName:  "Rogers",
			Email:     "frank@gmail.com",
		}, nil
	}
}
