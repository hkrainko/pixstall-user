package user

import (
	"context"
	authModel "pixstall-user/domain/auth/model"
	userModel "pixstall-user/domain/user/model"
)

type Repo interface {
	SaveAuthUser(ctx context.Context, authUserInfo *authModel.AuthUserInfo) (*userModel.User, error)
	UpdateUser(ctx context.Context, userID string, user *userModel.UserUpdater) error
	UpdateUserByAuthID(ctx context.Context, authID string, user *userModel.UserUpdater) error
	GetUserByAuthID(ctx context.Context, authID string) (*userModel.User, error)
	GetUser(ctx context.Context, userID string) (*userModel.User, error)
	GetUserDetails(ctx context.Context, userID string) (*userModel.User, error)
	IsUserExist(ctx context.Context, userID string) (*bool, error)
}