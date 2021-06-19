package model

import (
	model2 "pixstall-user/domain/user/model"
)

type ArtistUpdater struct {
	ArtistID    string
	UserName    *string
	ProfilePath *string
	Email       *string
	Birthday    *string
	Gender      *string
	State       *model2.UserState
}

func NewArtistUpdaterFromUserUpdater(updater model2.UserUpdater) ArtistUpdater {
	return ArtistUpdater{
		ArtistID:    updater.UserID,
		UserName:    updater.UserName,
		ProfilePath: updater.ProfilePath,
		Email:       updater.Email,
		Birthday:    updater.Birthday,
		Gender:      updater.Gender,
		State:       updater.State,
	}
}
