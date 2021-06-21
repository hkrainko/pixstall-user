package msg_broker

import (
	"context"
	model3 "pixstall-user/domain/artist/model"
	"pixstall-user/domain/commission/model"
	model2 "pixstall-user/domain/reg/model"
	userModel "pixstall-user/domain/user/model"
)

type Repo interface {
	SendCreateArtistCmd(ctx context.Context, info *model2.RegInfo) error
	SendUpdateArtistCmd(ctx context.Context, updater *model3.ArtistUpdater) error
	SendUserUpdatedEvent(ctx context.Context, updater *userModel.UserUpdater) error
	SendCommissionUserValidationEvent(ctx context.Context, usersValidation model.CommissionUsersValidation) error
	SendUpdateCommissionCmd(ctx context.Context, updater model.CommissionUpdater) error
}
