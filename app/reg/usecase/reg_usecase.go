package usecase

import (
	"context"
	"pixstall-user/app/domain/reg"
	"pixstall-user/app/domain/reg/model"
	"pixstall-user/app/domain/user"
	userModel "pixstall-user/app/domain/user/model"
)

type regUseCase struct {
	userRepo user.Repo
}

func NewRegUseCase(userRepo user.Repo) reg.UseCase {
	return &regUseCase{
		userRepo: userRepo,
	}
}

func (r regUseCase) Registration(ctx context.Context, info *model.RegInfo) error {
	updater := userModel.UserUpdater{
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

	err := r.userRepo.UpdateUserByAuthID(ctx, info.AuthID, &updater)
	if err != nil {
		return err
	}

	if info.RegAsArtist {
		//TODO: notify server
	}

	return nil
}