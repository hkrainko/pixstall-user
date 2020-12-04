package reg

import (
	"context"
	"pixstall-user/app/domain/reg/model"
)

type UseCase interface {
	Registration(ctx context.Context, info *model.RegInfo) error
}