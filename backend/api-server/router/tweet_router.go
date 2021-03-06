package router

import (
	"backend/api-server/controller"
	"backend/api-server/middleware"

	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/guregu/dynamo"
	"github.com/labstack/echo"
)

func TweetRouter(e *echo.Echo, db *dynamo.DB, auth *cognito.CognitoIdentityProvider, upload *s3manager.Uploader) {
	tweetController := controller.NewTweetController(db, auth, upload)
	r := e.Group("/tweets")
	r.Use(middleware.AuthMiddleware(auth))
	r.GET("", tweetController.TweetsIndex)
	r.POST("", tweetController.Post)
	r.GET("/:id", tweetController.Index)
	r.POST("/:id/likes", tweetController.Like)
	r.POST("/:id/retweets", tweetController.Retweet)
}
