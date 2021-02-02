package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	authModel "pixstall-user/domain/auth/model"
	commissionModel "pixstall-user/domain/commission/model"
	indexModel "pixstall-user/domain/inbox/model"
	"pixstall-user/domain/user/model"
	"time"
)

type User struct {
	ObjectID        primitive.ObjectID         `bson:"_id,omitempty"`
	UserID          string                     `bson:"userId,omitempty"`
	AuthID          string                     `bson:"authId,omitempty"`
	UserName        string                     `bson:"userName,omitempty"`
	AuthType        string                     `bson:"authType,omitempty"`
	Email           string                     `bson:"email,omitempty"`
	Birthday        string                     `bson:"birthday,omitempty"`
	Gender          string                     `bson:"gender,omitempty"`
	ProfilePath     string                     `bson:"profilePath,omitempty"`
	IsArtist        bool                       `bson:"isArtist"`
	State           model.UserState            `bson:"state,omitempty"`
	Inbox           indexModel.Inbox           `bson:"inbox,omitempty"`
	Commission      commissionModel.Commission `bson:"commission,omitempty"`
	RegTime         time.Time                  `bson:"regTime,omitempty"`
	LastUpdatedTime time.Time                  `bson:"lastUpdatedTime,omitempty"`
}

func (u *User) ToDomainUser() *model.User {
	return &model.User{
		UserID:          u.UserID,
		AuthID:          u.AuthID,
		UserName:        u.UserName,
		AuthType:        u.AuthType,
		Email:           u.Email,
		Birthday:        u.Birthday,
		Gender:          u.Gender,
		ProfilePath:     u.ProfilePath,
		IsArtist:        u.IsArtist,
		State:           u.State,
		Inbox:           u.Inbox,
		Commission:      u.Commission,
		RegTime:         u.RegTime,
		LastUpdatedTime: u.LastUpdatedTime,
	}
}

func NewFromAuthUserInfo(u *authModel.AuthUserInfo) *User {
	return &User{
		AuthID:          u.ID,
		AuthType:        u.AuthType,
		UserName:        u.UserName,
		Email:           u.Email,
		Birthday:        u.Birthday,
		Gender:          u.Gender,
		State:           model.UserStatePending,
		LastUpdatedTime: time.Now(),
	}
}
