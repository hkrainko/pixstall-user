package model

type UserUpdater struct {
	UserID      string
	UserName    string
	Email       string
	Birthday    string
	Gender      string
	ProfilePath string
	State       UserState
	IsArtist    *bool
}