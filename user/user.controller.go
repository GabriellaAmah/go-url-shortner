package user

import (
	"encoding/json"

	"net/http"

	"github.com/GabriellaAmah/go-url-shortner/util"
)

type userController struct {
}

func (uc userController) CreateUser(res http.ResponseWriter, request *http.Request) {
	var createUserData User

	err := json.NewDecoder(request.Body).Decode(&createUserData)
	if err != nil {
		util.HttpErrorResponse(res, util.MakeError{Code: 500, Message: "Internal server error", Error: err})
		return
	}

	data, serviceError := UserService.CreateUser(createUserData)
	if serviceError.IsError {
		util.HttpErrorResponse(res, serviceError)
		return
	}

	util.HttpSuccessResponse(res, 200, "User successfully created", data)

}

func (uc userController) LoginUser(res http.ResponseWriter, request *http.Request) {
	var loginUser LoginUser

	err := json.NewDecoder(request.Body).Decode(&loginUser)
	if err != nil {
		util.HttpErrorResponse(res, util.MakeError{Code: 500, Message: "Internal server error", Error: err})
		return
	}

	data, serviceError := UserService.LoginUser(loginUser)
	if serviceError.IsError {
		util.HttpErrorResponse(res, serviceError)
		return
	}

	util.HttpSuccessResponse(res, 200, "User successfully logged in", data)

}

var UserController = userController{}
