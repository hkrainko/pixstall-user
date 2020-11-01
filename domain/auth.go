package domain

import "context"

type AuthUsecase interface {
	GetAuthURL(ctx context.Context, authType string) (string, error)
}

type AuthRepository interface {
	GetAuthURL(ctx context.Context, authType string) (string, error)
}