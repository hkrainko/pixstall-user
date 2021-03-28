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
	AuthID          string                     `bson:"authId"`
	UserName        string                     `bson:"userName"`
	AuthType        string                     `bson:"authType"`
	Email           string                     `bson:"email"`
	Birthday        string                     `bson:"birthday"`
	Gender          string                     `bson:"gender"`
	ProfilePath     string                     `bson:"profilePath"`
	IsArtist        bool                       `bson:"isArtist"`
	State           model.UserState            `bson:"state"`
	Inbox           indexModel.Inbox           `bson:"inbox,omitempty"`
	Commission      commissionModel.Commission `bson:"commission,omitempty"`
	RegTime         time.Time                  `bson:"regTime"`
	LastUpdatedTime time.Time                  `bson:"lastUpdatedTime"`
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
