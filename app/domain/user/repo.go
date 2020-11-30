package user

import (
	"context"
	authModel "pixstall-user/app/domain/auth/model"
	userModel "pixstall-user/app/domain/user/model"
)

type Repo interface {
	SaveAuthUser(ctx context.Context, authUserInfo *authModel.AuthUserInfo) (*userModel.User, error)
	UpdateUser(ctx context.Context, user *userModel.User) error
	GetUser(ctx context.Context, userID string) (*userModel.User, error)
}