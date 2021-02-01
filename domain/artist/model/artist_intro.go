package model

type ArtistIntro struct {
	YearOfDrawing *int      `json:"yearOfDrawing" bson:"yearOfDrawing,omitempty"`
	ArtTypes      *[]string `json:"artTypes" bson:"artTypes,omitempty"`
}
