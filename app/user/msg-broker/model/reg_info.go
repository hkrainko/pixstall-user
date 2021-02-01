package model

import (
	"pixstall-user/domain/reg/model"
)

type RegInfo struct {
	*model.RegInfo
	ProfilePath string `json:"profilePath"`
}

func NewRegInfoFromDomainRegInfo(dRegInfo *model.RegInfo, ProfilePath string) *RegInfo {
	return &RegInfo{
		dRegInfo,
		ProfilePath,
	}
}

func (r *RegInfo) ToDomainUser() *model.RegInfo {
	return r.RegInfo
}
