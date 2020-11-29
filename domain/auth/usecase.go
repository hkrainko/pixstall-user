package auth

import (
	"context"
	authModel "pixstall-user/domain/auth/model"
)

type UseCase interface {
	HandleAuthCallback(ctx context.Context, authCallBack authModel.Callback) (*authModel.AuthUserInfo, error)
	GetAuthURL(ctx context.Context, authType string) (string, error)
}