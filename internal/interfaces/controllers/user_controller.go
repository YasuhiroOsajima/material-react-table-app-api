package controllers

import (
	"net/http"

	"github.com/YasuhiroOsajima/material-react-table-app-api/internal/usecase"
)

type UsersResult struct {
	Users any `json:"users"`
}

type UserController struct {
	UserInteractor *usecase.UserInteractor
}

func NewUserController(interactor *usecase.UserInteractor) *UserController {
	return &UserController{UserInteractor: interactor}
}

func (controller *UserController) GetUsers(ctx Context) {
	users := controller.UserInteractor.GetUsers()

	ctx.JSON(http.StatusOK, UsersResult{Users: users})
}
