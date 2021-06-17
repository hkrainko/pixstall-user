package user

import (
	"context"
	"pixstall-user/domain/file/model"
	domainUserModel "pixstall-user/domain/user/model"
)

type UseCase interface {
	GetUser(ctx context.Context, userID string) (*domainUserModel.User, error)
	GetAuthUser(ctx context.Context, userID string) (*domainUserModel.AuthUser, error)
	UpdateUser(ctx context.Context, updater *domainUserModel.UserUpdater, profile *model.ImageFile) (*string, error)
}
