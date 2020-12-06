package model

import (
	"pixstall-user/app/domain/reg/model"
)

type RegInfo struct {
	*model.RegInfo
}

func NewRegInfoFromDomainRegInfo(dRegInfo *model.RegInfo) *RegInfo {
	return &RegInfo{
		dRegInfo,
	}
}

func (r *RegInfo) ToDomainUser() *model.RegInfo {
	return r.RegInfo
}