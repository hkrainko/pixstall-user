package reg

import (
	"context"
	"image"
	"pixstall-user/domain/reg/model"
)

type UseCase interface {
	Registration(ctx context.Context, info *model.RegInfo, pngImage *image.Image) error
}