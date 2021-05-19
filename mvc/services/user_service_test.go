package services

import (
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
	"github.com/kirankothule/golang-microservice/mvc/domain"
	"github.com/kirankothule/golang-microservice/mvc/utils"
)

var(
	userDaoMock usersDaoMock
	getUserFunction func(userID int64) (*domain.User, *utils.ApplicationErr)
)

func init(){
	domain.UserDao=&usersDaoMock{}
}
type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userID int64) (*domain.User, *utils.ApplicationErr){
	return getUserFunction(userID)
}
func TestGetUserNotFoundInDatabase(t *testing.T){
	getUserFunction= func(userID int64) (*domain.User, *utils.ApplicationErr){
		return nil, &utils.ApplicationErr{
			StatusCode: http.StatusNotFound,
			Message: "User 0 not found",
		}
	}
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "User 0 not found", err.Message)
}

func TestGetUserFound(t *testing.T) {
	getUserFunction= func(userID int64) (*domain.User, *utils.ApplicationErr){
		return &domain.User{
			ID: 123,
		}, nil
	}
	user, err := UserService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
}