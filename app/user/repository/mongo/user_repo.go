package mongo

import (
	"context"
	"pixstall-user/app/domain/user"
	"pixstall-user/app/domain/user/model"
)

type mongoUserRepo struct {

}

func NewMongoUserRepo() user.Repo {
	return &mongoUserRepo{}
}

func (m mongoUserRepo) SaveUser(ctx context.Context, user *model.User) error {
	panic("implement me")
}

func (m mongoUserRepo) UpdateUser(ctx context.Context, user *model.User) error {
	panic("implement me")
}

func (m mongoUserRepo) GetUser(ctx context.Context, userID string) (*model.User, error) {
	panic("implement me")
}