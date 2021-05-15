package services

import (
	"github.com/kirankothule/golang-microservice/mvc/domain"
	"github.com/kirankothule/golang-microservice/mvc/utils"
)

type userSerivce struct {}

var (
	UserService userSerivce
)
// GetUser service will return user data from DAO
func (u userSerivce) GetUser(userID int64) (*domain.User, *utils.ApplicationErr) {
	return domain.UserDao.GetUser(userID)
}
