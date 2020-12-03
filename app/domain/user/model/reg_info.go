package model

type RegInfo struct {
	UserID        string     `json:"userId"`
	DisplayName   string     `json:"name"`
	Email         string     `json:"email"`
	Birthday      string     `json:"birthday"`
	Gender        string     `json:"gender"`
	RegAsArtist   bool       `json:"regAsArtist"`
	RegArtistInfo ArtistInfo `json:"regArtistInfo"`
}