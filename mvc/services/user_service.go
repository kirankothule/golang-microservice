package services

import (
	"github.com/kirankothule/golang-microservice/mvc/domain"
	"github.com/kirankothule/golang-microservice/mvc/utils"
)

// GetUser service will return user data from DAO
func GetUser(userID int64) (*domain.User, *utils.ApplicationErr) {
	return domain.GetUser(userID)
}
