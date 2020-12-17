package model

import (
	commissionModel "pixstall-user/domain/commission/model"
	indexModel "pixstall-user/domain/inbox/model"
)

type User struct {
	UserID      string
	AuthID      string
	UserName    string
	AuthType    string
	Token       string
	Email       string
	Birthday    string
	Gender      string
	ProfilePath string
	IsArtist    bool
	State       UserState
	Inbox       indexModel.Inbox
	Commission  commissionModel.Commission
}

func (u User) Error() string {
	panic("implement me")
}
