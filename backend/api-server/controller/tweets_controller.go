package controller

import (
	"backend/api-server/domain/entity"
	"backend/api-server/model"
	"net/http"
	"strconv"
	"strings"

	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/guregu/dynamo"
	"github.com/labstack/echo"
)

type TweetsController struct {
	tweetModel model.TweetModel
	userModel  model.UserModel
	tlModel    model.TimelineModel
}

func NewTweetController(db *dynamo.DB, auth *cognito.CognitoIdentityProvider) TweetsController {
	return TweetsController{
		tweetModel: model.NewTweetModel(db, auth),
		userModel:  model.NewUserModel(db, auth),
		tlModel:    model.NewTimelineModel(db),
	}
}

// ツイート取得
func (tc *TweetsController) Index(c echo.Context) error {
	id := c.Get("userID").(string)
	tweet := tc.tweetModel.Get(id)
	return c.JSON(http.StatusOK, tweet)
}

// ツイート一覧取得
func (tc *TweetsController) TweetsIndex(c echo.Context) error {
	userID := c.Get("userID").(string)
	tweets, err := tc.tlModel.Get(userID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, tweets)
}

// ツイート投稿
func (tc *TweetsController) Post(c echo.Context) error {
	userID := c.Get("userID").(string)
	currentUser, err := tc.userModel.Get(userID)
	if err != nil {
		return err
	}

	t := new(entity.PostTweet)
	c.Bind(t)
	content := strings.TrimSpace(t.Content)

	// TODO: エラーハンドリングの分離
	if content == "" {
		resp := CreateErrorMessage("Content is Required")
		return c.JSON(http.StatusBadRequest, resp)
	}
	if len(t.Content) > 420 {
		resp := CreateErrorMessage("Content is over 140")
		return c.JSON(http.StatusBadRequest, resp)
	}
	tc.tweetModel.Create(t, currentUser)

	// TODO: 作成されたtweetを返すようにしたい
	resp := CreateErrorMessage("POST success")
	return c.JSON(http.StatusOK, resp)
}

// いいね
func (tc *TweetsController) Like(c echo.Context) error {
	id := c.Get("userID").(string)
	tweetID := c.Param("id")
	t, _ := strconv.Atoi(tweetID)
	tc.tweetModel.Like(t, id)
	return c.String(http.StatusOK, "Like")
}

// リツイート
func (tc *TweetsController) Retweet(c echo.Context) error {
	id := c.Get("userID").(string)
	tweetID := c.Param("id")
	t, _ := strconv.Atoi(tweetID)
	tc.tweetModel.Retweet(t, id)
	return c.String(http.StatusOK, "Retweet")
}
