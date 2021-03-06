package router

import (
	"backend/api-server/controller"
	"backend/api-server/middleware"

	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/guregu/dynamo"
	"github.com/labstack/echo"
)

func UserRouter(e *echo.Echo, db *dynamo.DB, auth *cognito.CognitoIdentityProvider, upload *s3manager.Uploader) {
	userController := controller.NewUserController(db, auth, upload)

	// 認証なしのrouting
	authRouter := e.Group("/auth")
	authRouter.POST("/signup", userController.RegisterUser)
	authRouter.POST("/signin", userController.Signin)
	authRouter.POST("/refresh", userController.Refresh)

	// 認証ありのrouting
	userRouter := e.Group("/user")
	userRouter.Use(middleware.AuthMiddleware(auth))
	userRouter.GET("", userController.GetCurrentUser)
	userRouter.PUT("", userController.UpdateUser)
	userRouter.GET("/:userID", userController.Get)
	userRouter.GET("/:userID/tweets", userController.GetUserTL)
	userRouter.GET("/:userID/likes", userController.GetLikeTweets)
	userRouter.POST("/:followedUserID/follow", userController.Follow)
	userRouter.DELETE("/:followedUserID/follow", userController.Unfollow)
}
