package model

type ArtistInfo struct {
	YearOfDrawing int      `json:"yearOfDrawing"`
	ArtTypes      []string `json:"artTypes"`
}