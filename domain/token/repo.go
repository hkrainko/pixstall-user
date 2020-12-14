package token

import "context"

type Repo interface {
	GenerateAPIToken(ctx context.Context, userID string) (string, error)
	GenerateRegToken(ctx context.Context, userID string) (string, error)
	GenerateRefreshToken(ctx context.Context, userID string) (string, error)
}