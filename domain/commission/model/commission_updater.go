package model

type CommissionUpdater struct {
	OpenCommID           string  `json:"openCommissionId"`
	ArtistID             *string `json:"artistId"`
	ArtistName           *string `json:"artistName"`
	ArtistProfilePath    *string `json:"artistProfilePath"`
	RequesterName        *string `json:"requesterName"`
	RequesterProfilePath *string `json:"requesterProfilePath"`
}
