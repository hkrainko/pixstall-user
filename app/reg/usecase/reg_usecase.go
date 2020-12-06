package usecase

import (
	"context"
	"log"
	"pixstall-user/app/domain/reg"
	"pixstall-user/app/domain/reg/model"
	"pixstall-user/app/domain/user"
	userModel "pixstall-user/app/domain/user/model"
)

type regUseCase struct {
	userRepo   user.Repo
	artistRepo user.MsgBroker
}

func NewRegUseCase(userRepo user.Repo, artistRepo user.MsgBroker) reg.UseCase {
	return &regUseCase{
		userRepo: userRepo,
		artistRepo: artistRepo,
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
		err := r.artistRepo.SendRegisterArtistMsg(ctx, info)
		//not return err
		if err != nil {
			log.Printf("SendRegisterArtistMsg err %v", err)
		}
	} else {
		err := r.artistRepo.SendRegisterUserMsg(ctx, info)
		//not return err
		if err != nil {
			log.Printf("SendRegisterUserMsg err %v", err)
		}
	}

	return nil
}