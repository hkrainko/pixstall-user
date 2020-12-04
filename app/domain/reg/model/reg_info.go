package model

import "pixstall-user/app/domain/user/model"

type RegInfo struct {
	AuthID        string           `json:"authId"`
	UserID        string           `json:"userId"`
	DisplayName   string           `json:"name"`
	Email         string           `json:"email"`
	Birthday      string           `json:"birthday"`
	Gender        string           `json:"gender"`
	RegAsArtist   bool             `json:"regAsArtist"`
	RegArtistInfo model.ArtistInfo `json:"regArtistInfo"`
}
