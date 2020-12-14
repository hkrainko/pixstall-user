package token

import "context"

type Repo interface {
	GenerateToken(ctx context.Context, userID string) (string, error)
}