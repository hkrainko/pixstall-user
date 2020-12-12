package image

import (
	"context"
	"image"
)

type Repo interface {
	SaveImage(ctx context.Context, path string, fileName string, image image.Image) error
}