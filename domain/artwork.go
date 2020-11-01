package domain

import "time"

type Artwork struct {
	ID            string    `json:"id"`
	TrustorID     string    `json:"trustorId"`
	ArtistID      string    `json:"artistId"`
	CommissionDT  time.Time `json:"commissionDt"`
	CompletedDate time.Time `json:"completedDt"`
	ArtworkType   string    `json:"artworkType"`
	FileType      string    `json:"fileType"`
	FileSize      float64   `json:"fileSize"`
	Width         float64   `json:"width"`
	Height        float64   `json:"height"`
	Title         string    `json:"title"`
	Body          string    `json:"body"`
	Tabs          []string  `json:"tabs"`
	LikerIDs      []string  `json:"likerIds"`
	ViewerNum     float64   `json:"viewerNum"`
}

type ArtworkUsecase interface {
	FindByID(id string) (*Artwork, error)
}

type ArtworkRepository interface {
	FindByID(id string) (*Artwork, error)
}