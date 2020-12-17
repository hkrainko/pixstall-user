package usecase

import (
	"context"
	"pixstall-user/domain/image"
	"pixstall-user/domain/user"
	"pixstall-user/domain/user/model"
)

type userUseCase struct {
	userRepo user.Repo
	imageRepo image.Repo
}

func NewUserUseCase(userRepo user.Repo, imageRepo image.Repo) user.UseCase {
	return &userUseCase{
		userRepo: userRepo,
		imageRepo: imageRepo,
	}
}

func (u userUseCase) GetUser(ctx context.Context, userID string) (*model.User, error) {
	dUser, err := u.userRepo.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return dUser, nil
}

func (u userUseCase) GetUserDetails(ctx context.Context, userID string) (*model.User, error) {
	dUser, err := u.userRepo.GetUserDetails(ctx, userID)
	if err != nil {
		return nil, err
	}
	return dUser, nil
}