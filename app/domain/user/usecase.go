package user

import (
	"context"
	userModel "pixstall-user/app/domain/user/model"
)

type UseCase interface {
	CompleteRegistration(ctx context.Context, info *userModel.RegInfo) error
}