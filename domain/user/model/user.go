package model

import "pixstall-user/domain/artist/model"

type User struct {
	UserID     string
	AuthID     string
	UserName   string
	AuthType   string
	Email      string
	Birthday   string
	Gender     string
	PhotoURL   string
	IsArtist   bool
	ArtistInfo model.ArtistIntro
	State      UserState
}

func (u User) Error() string {
	panic("implement me")
}
