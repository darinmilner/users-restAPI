package controllers

import (
	"encoding/json"
	"golang/microservices/services"
	"golang/microservices/utils"
	"log"
	"net/http"
	"strconv"
)

func GetUser(res http.ResponseWriter, req *http.Request) {

	log.Print("User Route")
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	log.Println(userId)
	if err != nil {
		//Return bad request to client
		apiErr := &utils.ApplicationError{
			Message:    "User id should be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "Bad Request",
		}

		jsonVal, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write([]byte(jsonVal))
		return
	}

	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil {
		//handle error
		jsonVal, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write([]byte(jsonVal))
		return
	}

	//return user to client
	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
