package usecase

import (
	"context"
	"pixstall-user/domain/auth"
	authModel "pixstall-user/domain/auth/model"
)

type authUseCase struct {
	authRepo auth.Repo
}

func NewAuthUseCase(repo auth.Repo) auth.UseCase {
	return &authUseCase{
		authRepo: repo,
	}
}

func (a authUseCase) GetAuthURL(ctx context.Context, authType string) (string, error) {
	return a.authRepo.GetAuthURL(ctx, authType)
}

func (a authUseCase) HandleAuthCallback(ctx context.Context, authCallBack authModel.Callback) (*authModel.AuthUserInfo, error) {
	userInfo, err := a.authRepo.GetAuthorizedUserInfo(ctx, authCallBack)
	if err != nil {
		return nil, err
	}
	//TODO: Event for new user
	return userInfo, nil
}
