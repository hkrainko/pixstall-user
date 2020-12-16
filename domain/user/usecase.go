package user

import (
	"context"
	domainUserModel "pixstall-user/domain/user/model"
)

type UseCase interface {
	GetUser(ctx context.Context, userID string) (*domainUserModel.User, error)
}