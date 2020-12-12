package usecase

import (
	"pixstall-user/domain/image"
	"pixstall-user/domain/user"
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
