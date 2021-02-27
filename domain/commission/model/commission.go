package model

type Commission struct {
	ID          string `json:"id"`
	ArtistID    string `json:"artistID"`
	RequesterID string `json:"requesterID"`
}
