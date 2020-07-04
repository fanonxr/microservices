package controllers

import (
	"github.com/gin-gonic/gin"
	"microservices/mvc/services"
	"microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		// creating our app error
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be an number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		// responding in json form
		utils.RespondError(c, apiErr)
		return
	}

	// calling the user service
	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}
	// handling the case where we have a valid user
	utils.Respond(c, http.StatusOK, user)
}