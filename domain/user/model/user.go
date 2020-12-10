package model

type User struct {
	UserID     string
	AuthID     string
	UserName   string
	AuthType   string
	Email      string
	Birthday   string
	Gender     string
	ProfilePath   string
	IsArtist   bool
	State      UserState
}

func (u User) Error() string {
	panic("implement me")
}
