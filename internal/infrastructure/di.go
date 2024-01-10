package infrastructure

import (
	"github.com/YasuhiroOsajima/material-react-table-app-api/internal/interfaces/controllers"
	"github.com/YasuhiroOsajima/material-react-table-app-api/internal/usecase"
)

var UserInteractor *usecase.UserInteractor
var userCtrl *controllers.UserController

func init() {
	UserInteractor = usecase.NewUserInteractor()
	userCtrl = controllers.NewUserController(UserInteractor)
}
