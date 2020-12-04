package model

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
	ArtistInfo ArtistInfo
	State      UserState
}

func (u User) Error() string {
	panic("implement me")
}
