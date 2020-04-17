package model

import (
	"backend/api-server/domain/entity"
	"fmt"

	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/guregu/dynamo"
)

type UserModel struct {
	userTable dynamo.Table
	auth      *cognito.CognitoIdentityProvider
}

func NewUserModel(db *dynamo.DB, auth *cognito.CognitoIdentityProvider) UserModel {
	return UserModel{
		userTable: db.Table("Users"),
		auth:      auth,
	}
}

func (um *UserModel) All() *[]entity.User {
	return nil
}

func (um *UserModel) Get(id string) (*entity.User, error) {
	user := new(entity.User)
	if err := um.userTable.Get("id", id).One(user); err != nil {
		fmt.Println(err)
		return user, err
	}
	return user, nil
}

func (um *UserModel) Update(t *entity.User) {
	if err := um.userTable.Put(t).Run(); err != nil {
		fmt.Println(err)
	}
	return
}

func (um *UserModel) GetOrCreateDummyUser() string {
	dummyID := "dummy"
	user, err := um.Get(dummyID)
	if err == nil {
		return user.ID
	}
	newUser := entity.User{
		ID:         dummyID,
		ScreenName: "anonymous",
		IconUrl:    "https://pbs.twimg.com/profile_images/1136178449779810304/1e0ghs3t_400x400.jpg",
	}
	if err := um.userTable.Put(newUser).Run(); err != nil {
		fmt.Println(err)
	}
	return dummyID
}

func (um *UserModel) Regist(u *entity.SignUpUser) entity.User {
	user := entity.User{
		ID:         u.ID,
		ScreenName: u.ScreenName,
	}
	if err := um.userTable.Put(user).Run(); err != nil {
		fmt.Println(err)
	}
	return user
}

func (um *UserModel) Follow(userID string, followedID string) (*entity.User, *entity.User) {
	userInfo := new(entity.User)
	followedUserInfo := new(entity.User)
	if err := um.userTable.Get("id", userID).One(&userInfo); err != nil {
		fmt.Println(err)
	}
	if err := um.userTable.Get("id", followedID).One(&followedUserInfo); err != nil {
		fmt.Println(err)
	}
	userInfo.FollowIDs = append(userInfo.FollowIDs, followedID)
	followedUserInfo.FollowedIDs = append(followedUserInfo.FollowedIDs, userID)
	return userInfo, followedUserInfo
}

func (um *UserModel) UnFollow(userID string, followedID string) (*entity.User, *entity.User) {
	userInfo := new(entity.User)
	followedUserInfo := new(entity.User)
	if err := um.userTable.Get("id", userID).One(&userInfo); err != nil {
		fmt.Println(err)
	}
	if err := um.userTable.Get("id", followedID).One(&followedUserInfo); err != nil {
		fmt.Println(err)
	}
	userInfo.FollowIDs = RemoveUser(userInfo.FollowIDs, followedID)
	followedUserInfo.FollowedIDs = RemoveUser(followedUserInfo.FollowedIDs, userID)
	return userInfo, followedUserInfo
}

func RemoveUser(userIDList []string, userID string) []string {
	list := []string{}
	for _, v := range userIDList {
		if v == userID {
			continue
		}
		list = append(list, v)
	}
	return list
}
