package controller

import (
	"backend/api-server/model"
	"net/http"

	"github.com/guregu/dynamo"
	"github.com/labstack/echo"
)

type UsersController struct {
	userModel model.UserModel
}

func NewUserController(db *dynamo.DB) UsersController {
	return UsersController{
		userModel: model.NewUserModel(db),
	}
}

func (uc *UsersController) UserIndex(c echo.Context) error {
	uc.userModel.All()
	return c.String(http.StatusOK, "OK!")
}
