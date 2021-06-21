package msg

import (
	"pixstall-user/domain/reg/model"
)

type RegInfo struct {
	model.RegInfo `json:",inline"`
}

func NewCreateArtistCmdMsg(dRegInfo model.RegInfo) RegInfo {
	return RegInfo{
		dRegInfo,
	}
}
