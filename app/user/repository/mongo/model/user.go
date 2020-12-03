package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	authModel "pixstall-user/app/domain/auth/model"
	"pixstall-user/app/domain/user/model"
)

type User struct {
	ObjectID primitive.ObjectID `bson:"_id,omitempty"`
	UserID   string             `bson:"userId,omitempty"`
	AuthID   string             `bson:"authId,omitempty"`
	UserName string             `bson:"userName,omitempty"`
	AuthType string             `bson:"authType,omitempty"`
	Token    string             `bson:"token,omitempty"`
	Email    string             `bson:"emil,omitempty"`
	Birthday string             `bson:"birthday,omitempty"`
	Gender   string             `bson:"gender,omitempty"`
	PhotoURL string             `bson:"photoUrl,omitempty"`
	State    model.UserState    `bson:"state,omitempty"`
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
		PhotoURL: u.PhotoURL,
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
		PhotoURL: u.PhotoURL,
		State:    "P",
	}
}
