package usecase

import (
	"context"
	"log"
	"pixstall-user/domain/image"
	"pixstall-user/domain/reg"
	"pixstall-user/domain/reg/model"
	"pixstall-user/domain/user"
	userModel "pixstall-user/domain/user/model"
)

type regUseCase struct {
	userRepo      user.Repo
	userMsgBroker user.MsgBroker
	imageRepo image.Repo
}

func NewRegUseCase(userRepo user.Repo, userMsgBroker user.MsgBroker, imageRepo image.Repo) reg.UseCase {
	return &regUseCase{
		userRepo:      userRepo,
		userMsgBroker: userMsgBroker,
		imageRepo: imageRepo,
	}
}

func (r regUseCase) Registration(ctx context.Context, info *model.RegInfo) error {
	updater := userModel.UserUpdater{
		UserID:     info.UserID,
		UserName:   info.DisplayName,
		Email:      info.Email,
		Birthday:   info.Birthday,
		Gender:     info.Gender,
		ProfilePath:   "",
		State:      "A",
		IsArtist:   &info.RegAsArtist,
	}

	err := r.userRepo.UpdateUserByAuthID(ctx, info.AuthID, &updater)
	if err != nil {
		return err
	}

	if info.RegAsArtist {
		err := r.userMsgBroker.SendRegisterArtistMsg(ctx, info)
		//not return err
		if err != nil {
			log.Printf("SendRegisterArtistMsg err %v", err)
		}
	} else {
		err := r.userMsgBroker.SendRegisterUserMsg(ctx, info)
		//not return err
		if err != nil {
			log.Printf("SendRegisterUserMsg err %v", err)
		}
	}

	return nil
}
