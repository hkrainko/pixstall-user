package user

import (
	"context"
	domainUserModel "pixstall-user/domain/user/model"
)

type UseCase interface {
	GetUser(ctx context.Context, userID string) (*domainUserModel.User, error)
	GetUserDetails(ctx context.Context, userID string) (*domainUserModel.UserDetails, error)
	UpdateUser(ctx context.Context, updater *domainUserModel.UserUpdater) (*domainUserModel.User, error)
}
