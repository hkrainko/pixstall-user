package reg

import (
	"context"
	model2 "pixstall-user/domain/file/model"
	"pixstall-user/domain/reg/model"
	dUserModel "pixstall-user/domain/user/model"
)

type UseCase interface {
	Registration(ctx context.Context, info model.RegInfo, profile *model2.ImageFile) (authUser *dUserModel.AuthUser, err error)
}
