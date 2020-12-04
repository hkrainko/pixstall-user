package usecase

import (
	"context"
	"pixstall-user/app/domain/user"
	"pixstall-user/app/domain/user/model"
)

type userUseCase struct {
	userRepo user.Repo
}

func NewUserUseCase(userRepo user.Repo) user.UseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (u userUseCase) CompleteRegistration(ctx context.Context, info *model.RegInfo) error {

	updater := model.UserUpdater{
		UserID:     info.UserID,
		UserName:   info.DisplayName,
		Email:      info.Email,
		Birthday:   info.Birthday,
		Gender:     info.Gender,
		PhotoURL:   "",
		State:      "A",
		IsArtist:   &info.RegAsArtist,
		ArtistInfo: &info.RegArtistInfo,
	}

	err := u.userRepo.UpdateUserByAuthID(ctx, info.AuthID, &updater)
	if err != nil {
		return err
	}

	if info.RegAsArtist {
		//TODO: notify server
	}

	return nil
}
