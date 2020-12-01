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
	AuthType string             `bson:"authType,omitempty"`
	Token    string             `bson:"token,omitempty"`
	Email    string             `bson:"emil,omitempty"`
	Birthday string             `bson:"birthday,omitempty"`
	Gender   string             `bson:"gender,omitempty"`
	PhotoURL string             `bson:"photoUrl,omitempty"`
	State    model.UserState    `bson:"state,omitempty"`
}

func (u *User) toDomainUser() *model.User {
	return &model.User{
		UserID:   u.UserID,
		AuthID:   u.AuthID,
		AuthType: u.AuthType,
		Email:    u.Email,
		Birthday: u.Birthday,
		Gender:   u.Gender,
		PhotoURL: u.PhotoURL,
		State:    u.State,
	}
}

func NewFromUser(u *model.User) *User {
	return &User{
		UserID:   u.UserID,
		AuthID:   u.AuthID,
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
		UserID:   "",
		AuthID:   u.ID,
		AuthType: u.AuthType,
		Token:    u.Token,
		Email:    u.Email,
		Birthday: u.Birthday,
		Gender:   u.Gender,
		PhotoURL: u.PhotoURL,
		State:    "P",
	}
}
