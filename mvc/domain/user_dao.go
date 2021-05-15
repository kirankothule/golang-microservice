package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kirankothule/golang-microservice/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Lee", LastName: "temp", Email: "temp@gmail.com"},
	}
)

// GetUser Dao to get user from data base
func GetUser(userID int64) (*User, *utils.ApplicationErr) {

	if user := users[userID]; user != nil {
		log.Printf("User Found with id: %v", userID)
		return user, nil
	}

	log.Printf("User Not Found with id: %v", userID)

	return nil, &utils.ApplicationErr{
		Message:    fmt.Sprintf("User %v not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
