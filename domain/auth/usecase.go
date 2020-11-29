package auth

import (
	"context"
	authModel "pixstall-user/domain/auth/model"
)

type UseCase interface {
	GetAuthURL(ctx context.Context, authType string) (string, error)
	HandleAuthCallback(ctx context.Context, authCallBack authModel.Callback) (*authModel.AuthUserInfo, error)
}