package token

import "context"

type Repo interface {
	GenerateAuthToken(ctx context.Context, userID string) (string, error)
	GenerateRefreshToken(ctx context.Context, userID string) (string, error)
}