package controllers

import (
	"encoding/json"
	"microservices/mvc/services"
	"microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		appErr := &utils.ApplicationError{
			Message:    "user_id must be an number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(appErr)
		res.WriteHeader(appErr.StatusCode)
		res.Write(jsonValue)
		// just return the bad request to the client
		return
	}

	// calling the user service
	user, apiErr := services.UserService.GetUser(userId)

	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		// handle the error and return to the client
		res.WriteHeader(apiErr.StatusCode)
		res.Write([]byte(jsonValue))
		return
	}
	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}