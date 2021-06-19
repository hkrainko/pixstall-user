package msg

import (
	domainArtistModel "pixstall-user/domain/artist/model"
)

type ArtistUpdater struct {
	*domainArtistModel.ArtistUpdater
}

func NewArtistUpdaterFromDomainUserUpdater(u *domainArtistModel.ArtistUpdater) *ArtistUpdater {
	return &ArtistUpdater{
		u,
	}
}