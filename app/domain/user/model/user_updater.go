package model

type UserUpdater struct {
	UserName   string
	Email      string
	Birthday   string
	Gender     string
	PhotoURL   string
	State      UserState
	IsArtist   bool
	ArtistInfo ArtistInfo
}
