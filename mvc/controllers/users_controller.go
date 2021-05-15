package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/kirankothule/golang-microservice/mvc/services"
	"github.com/kirankothule/golang-microservice/mvc/utils"
)

// GetUser Controller function for user request
func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIDParam := req.URL.Query().Get("user_id")
	log.Printf("about to process user_id: %v", userIDParam)
	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		//return error as bad request
		apiErr := &utils.ApplicationErr{
			Message:    "User Id must be the number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}

	user, apiErr := services.GetUser(userID)

	if apiErr != nil {
		//Handle the error
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
