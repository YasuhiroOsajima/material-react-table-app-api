package controllers

import (
	"fmt"
	"net/http"

	"github.com/YasuhiroOsajima/material-react-table-app-api/internal/usecase"
)

type UsersResult struct {
	Users any `json:"users"`
}

type DeleteResult struct {
	Message string `json:"message"`
}

type ErrorResult struct {
	Error string `json:"error"`
}

type UserController struct {
	UserInteractor *usecase.UserInteractor
}

func NewUserController(interactor *usecase.UserInteractor) *UserController {
	return &UserController{UserInteractor: interactor}
}

func (controller *UserController) GetUsers(ctx Context) {
	users, err := controller.UserInteractor.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResult{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, UsersResult{Users: users})
}

func (controller *UserController) DeleteUser(ctx Context) {
	username := ctx.Param("username")

	err := controller.UserInteractor.DeleteUser(username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResult{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, DeleteResult{Message: fmt.Sprintf("user '%s' deleted", username)})
}
