package model

import (
	commissionModel "pixstall-user/domain/commission/model"
	indexModel "pixstall-user/domain/inbox/model"
	"time"
)

type User struct {
	UserID          string    `json:"userId"`
	AuthID          string    `json:"authId,omitempty"`
	UserName        string    `json:"userName"`
	AuthType        string    `json:"authType,omitempty"`
	Email           string    `json:"email"`
	Birthday        string    `json:"birthday"`
	Gender          string    `json:"gender"`
	ProfilePath     string    `json:"profilePath"`
	IsArtist        bool      `json:"isArtist"`
	State           UserState `json:"state"`
	Inbox           indexModel.Inbox
	Commission      commissionModel.Commission
	RegTime         time.Time `json:"regTime"`
	LastUpdatedTime time.Time
}

func (u User) Error() string {
	panic("implement me")
}
