package user

import (
	"context"
	domainRegModel "pixstall-user/domain/reg/model"
	domainUserModel "pixstall-user/domain/user/model"
)

type MsgBroker interface {
	SendRegisterArtistMsg(ctx context.Context, info *domainRegModel.RegInfo, profilePath string) error
	SendArtistUpdateMsg(ctx context.Context, updater *domainUserModel.UserUpdater) error
	SendRegisterUserMsg(ctx context.Context, info *domainRegModel.RegInfo, profilePath string) error
	SendUserUpdateMsg(ctx context.Context, updater *domainUserModel.UserUpdater) error
}
