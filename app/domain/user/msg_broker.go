package user

import (
	"context"
	domainRegModel "pixstall-user/app/domain/reg/model"
	domainUserModel "pixstall-user/app/domain/user/model"
)

type MsgBroker interface {
	SendRegisterArtistMsg(ctx context.Context, info *domainRegModel.RegInfo) error
	SendArtistUpdateMsg(ctx context.Context, updater *domainUserModel.UserUpdater) error
	SendRegisterUserMsg(ctx context.Context, info *domainRegModel.RegInfo) error
	SendUserUpdateMsg(ctx context.Context, updater *domainUserModel.UserUpdater) error
}
