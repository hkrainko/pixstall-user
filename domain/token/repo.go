package token

import "context"

type Repo interface {
	GenerateRegToken(ctx context.Context, authID string) (string, error)
	GenerateAPIToken(ctx context.Context, userID string) (string, error)
}