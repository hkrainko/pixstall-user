package msg_broker

import (
	"context"
	model3 "pixstall-user/domain/artist/model"
	"pixstall-user/domain/commission/model"
	model2 "pixstall-user/domain/reg/model"
)

type Repo interface {
	SendRegisterArtistMsg(ctx context.Context, info *model2.RegInfo) error
	SendArtistUpdateMsg(ctx context.Context, updater *model3.ArtistUpdater) error
	SendCommissionUserValidationMsg(ctx context.Context, usersValidation model.CommissionUsersValidation) error
	SendCommissionUpdateMsg(ctx context.Context, updater model.CommissionUpdater) error
}
