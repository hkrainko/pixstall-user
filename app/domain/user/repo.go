package user

import (
	"context"
	"pixstall-user/app/domain/user/model"
)

type Repo interface {
	SaveUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, userID string) (*model.User, error)
}