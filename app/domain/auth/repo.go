package auth

import (
	"context"
	authModel "pixstall-user/app/domain/auth/model"
)

type Repo interface {
	GetAuthURL(ctx context.Context, authType string) (string, error)
	GetAuthorizedUserInfo(ctx context.Context, authCallBack authModel.AuthCallback) (*authModel.AuthUserInfo, error)
}