package usecase

import (
	"context"
	"pixstall_server/app/domain"
)

type authUsecase struct {
	authRepo domain.AuthRepository
}

func NewAuthUsecase(repo domain.AuthRepository) domain.AuthUsecase {
	return &authUsecase{
		authRepo: repo,
	}
}

func (a authUsecase) GetAuthURL(ctx context.Context, authType string) (string, error) {
	return a.authRepo.GetAuthURL(ctx, authType)
}
