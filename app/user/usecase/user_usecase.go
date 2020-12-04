package usecase

import (
	"pixstall-user/app/domain/user"
)

type userUseCase struct {
	userRepo user.Repo
}

func NewUserUseCase(userRepo user.Repo) user.UseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}
