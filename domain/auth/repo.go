package auth

import (
	"context"
)

type Repo interface {
	GetAuthorizedUserInfo(ctx context.Context, authCallBack string) (*string, error)
	GetAuthURL(ctx context.Context, authType string) (string, error)
}