package auth

import "context"

type UseCase interface {
	HandleAuthCallBack(ctx context.Context, authCallBack string) (*string, error)
	GetAuthURL(ctx context.Context, authType string) (string, error)
}