package reg

import (
	"context"
	"image"
	"pixstall-user/domain/reg/model"
	dUserModel "pixstall-user/domain/user/model"
)

type UseCase interface {
	Registration(ctx context.Context, info model.RegInfo, pngImage image.Image) (authUser *dUserModel.AuthUser, err error)
}
