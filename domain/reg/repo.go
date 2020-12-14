package reg

import "context"

type Repo interface {
	GenerateRegToken(ctx context.Context, authID string) (string, error)
	ValidateRegToken(ctx context.Context, token string) (string, error)
}