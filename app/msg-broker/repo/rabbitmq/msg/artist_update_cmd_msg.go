package msg

import (
	domainArtistModel "pixstall-user/domain/artist/model"
)

type ArtistUpdateCmdMsg struct {
	domainArtistModel.ArtistUpdater `json:",inline"`
}

func NewUpdateArtistCmdMsg(u domainArtistModel.ArtistUpdater) ArtistUpdateCmdMsg {
	return ArtistUpdateCmdMsg{
		u,
	}
}