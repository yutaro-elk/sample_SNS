package controller

import (
	"backend/api-server/domain/entity"
	"backend/api-server/domain/services"
	"backend/api-server/model"
	"net/http"

	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/guregu/dynamo"
	"github.com/labstack/echo"
)

type UsersController struct {
	userModel   model.UserModel
	userService services.UserServices
}

func NewUserController(db *dynamo.DB, auth *cognito.CognitoIdentityProvider) UsersController {
	return UsersController{
		userModel:   model.NewUserModel(db, auth),
		userService: services.NewUserServices(auth),
	}
}

// ユーザー情報の取得
func (uc *UsersController) UserIndex(c echo.Context) error {
	userID := c.Param("userName")
	user, _ := uc.userModel.Get(userID)
	return c.JSON(http.StatusOK, user)
}

// フォローの取得
func (uc *UsersController) FollowsIndex(c echo.Context) error {
	userID := c.Param("userName")
	uc.userModel.All()
	return c.String(http.StatusOK, "GetFollows"+userID)
}

// フォロワーの取得
func (uc *UsersController) FollowersIndex(c echo.Context) error {
	userID := c.Param("userName")
	uc.userModel.All()
	return c.String(http.StatusOK, "GetFollowers"+userID)
}

// ユーザー情報更新
func (uc *UsersController) UpdateUser(c echo.Context) error {
	userID := c.Param("userName")
	uc.userModel.All()
	return c.String(http.StatusOK, "GetFollowers"+userID)
}

// ユーザー登録
func (uc *UsersController) RegisterUser(c echo.Context) error {
	u := new(entity.SignUpUser)
	c.Bind(u)

	if err := uc.userService.CreateUserOnCognito(u); err != nil {
		return c.JSON(http.StatusBadRequest, CreateErrorMessage(err.Error()))
	}

	resp := uc.userModel.Regist(u)

	return c.JSON(http.StatusCreated, resp)
}

// フォロー処理
func (uc *UsersController) Follow(c echo.Context) error {
	userID := c.Param("userName")
	uc.userModel.All()
	return c.String(http.StatusOK, "Follow"+userID)
}

// アンフォロー処理
func (uc *UsersController) Unfollow(c echo.Context) error {
	userID := c.Param("userName")
	uc.userModel.All()
	return c.String(http.StatusOK, "Unfollow"+userID)
}
