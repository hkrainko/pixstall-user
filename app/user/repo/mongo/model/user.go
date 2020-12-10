package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	authModel "pixstall-user/domain/auth/model"
	"pixstall-user/domain/user/model"
)

type User struct {
	ObjectID   primitive.ObjectID `bson:"_id,omitempty"`
	UserID     string             `bson:"userId,omitempty"`
	AuthID     string             `bson:"authId,omitempty"`
	UserName   string             `bson:"userName,omitempty"`
	AuthType   string             `bson:"authType,omitempty"`
	Token      string             `bson:"token,omitempty"`
	Email      string             `bson:"email,omitempty"`
	Birthday   string             `bson:"birthday,omitempty"`
	Gender     string             `bson:"gender,omitempty"`
	ProfilePath   string             `bson:"photoUrl,omitempty"`
	IsArtist   bool               `bson:"isArtist"`
	State      model.UserState    `bson:"state,omitempty"`
}

func (u *User) ToDomainUser() *model.User {
	return &model.User{
		UserID:   u.UserID,
		AuthID:   u.AuthID,
		UserName: u.UserName,
		AuthType: u.AuthType,
		Email:    u.Email,
		Birthday: u.Birthday,
		Gender:   u.Gender,
		ProfilePath: u.ProfilePath,
		IsArtist: u.IsArtist,
		State:    u.State,
	}
}

func NewFromAuthUserInfo(u *authModel.AuthUserInfo) *User {
	return &User{
		AuthID:   u.ID,
		AuthType: u.AuthType,
		UserName: u.UserName,
		Email:    u.Email,
		Birthday: u.Birthday,
		Gender:   u.Gender,
		State:    "P",
	}
}
