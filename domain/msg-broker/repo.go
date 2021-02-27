package msg_broker

import (
	"context"
	"pixstall-user/domain/commission/model"
	model2 "pixstall-user/domain/reg/model"
	domainUserModel "pixstall-user/domain/user/model"
)

type Repo interface {
	SendRegisterArtistMsg(ctx context.Context, info *model2.RegInfo) error
	SendArtistUpdateMsg(ctx context.Context, updater *domainUserModel.UserUpdater) error
	SendCommissionUserValidationMsg(ctx context.Context, usersValidation model.CommissionUsersValidation) error
	SendCommissionUpdateMsg(ctx context.Context, updater model.CommissionUpdater) error
}
