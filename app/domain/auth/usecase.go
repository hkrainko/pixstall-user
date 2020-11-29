package auth

import (
	"context"
	authModel "pixstall-user/app/domain/auth/model"
	"pixstall-user/app/domain/user/model"
)

type UseCase interface {
	GetAuthURL(ctx context.Context, authType string) (string, error)
	HandleAuthCallback(ctx context.Context, authCallBack authModel.AuthCallback) (*model.User, error)
}