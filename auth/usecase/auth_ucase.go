package usecase

import (
	"context"
	"pixstall-user/domain/auth"
)

type authUseCase struct {
	authRepo auth.Repo
}

func NewAuthUsecase(repo auth.Repo) auth.UseCase {
	return &auth.UseCase{
		authRepo: repo,
	}
}

func (a authUseCase) GetAuthURL(ctx context.Context, authType string) (string, error) {
	return a.authRepo.GetAuthURL(ctx, authType)
}
