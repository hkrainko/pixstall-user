package model

type UserUpdater struct {
	UserID     string
	UserName   string
	Email      string
	Birthday   string
	Gender     string
	PhotoURL   string
	State      UserState
	IsArtist   *bool
	ArtistInfo *ArtistInfo
}
