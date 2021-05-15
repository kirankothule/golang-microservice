package controllers

import (
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/kirankothule/golang-microservice/mvc/services"
	"github.com/kirankothule/golang-microservice/mvc/utils"
)

// GetUser Controller function for user request
func GetUser(c *gin.Context) {
	userIDParam := c.Param("user_id")
	log.Printf("about to process user_id: %v", userIDParam)
	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		//return error as bad request
		apiErr := &utils.ApplicationErr{
			Message:    "User Id must be the number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondErr(c,apiErr)
		return
	}

	user, apiErr := services.UserService.GetUser(userID)

	if apiErr != nil {
		//Handle the error
		utils.RespondErr(c,apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, user)
}
