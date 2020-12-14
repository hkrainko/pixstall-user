package usecase

import (
	"context"
	"github.com/google/uuid"
	"image"
	"log"
	"pixstall-user/app/utils"
	domainImage "pixstall-user/domain/image"
	"pixstall-user/domain/reg"
	"pixstall-user/domain/reg/model"
	"pixstall-user/domain/user"
	userModel "pixstall-user/domain/user/model"
	"strings"
)

type regUseCase struct {
	userRepo      user.Repo
	userMsgBroker user.MsgBroker
	imageRepo     domainImage.Repo
	regRepo       reg.Repo
}

func NewRegUseCase(userRepo user.Repo, userMsgBroker user.MsgBroker, imageRepo domainImage.Repo, regRepo reg.Repo) reg.UseCase {
	return &regUseCase{
		userRepo:      userRepo,
		userMsgBroker: userMsgBroker,
		imageRepo:     imageRepo,
		regRepo:       regRepo,
	}
}

func (r regUseCase) Registration(ctx context.Context, info *model.RegInfo, pngImage image.Image) error {

	//Check if authUser exist and in pending state
	extUser, err := r.userRepo.GetUserByAuthID(ctx, info.AuthID)
	if err != nil {
		return err
	}
	if extUser.State != userModel.UserStatePending {
		return userModel.UserErrorAuthIDAlreadyRegister
	}

	//Check if userId exist
	exist, err := r.userRepo.IsUserExist(ctx, info.UserID)
	if err != nil {
		return err
	}
	if *exist {
		return userModel.UserErrorDuplicateUser
	}

	//Upload image
	profilePath := func() string {
		if pngImage == nil {
			return ""
		}
		newUUID, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			return ""
		}
		fileName := newUUID.String()
		fileName = strings.ReplaceAll(fileName, "-", "")
		fileName = info.UserID + "_" + fileName
		//TODO: put profile path into other place
		path := "profile/"
		err = r.imageRepo.SaveImage(ctx, path, fileName+"_50", utils.ResizeImage(pngImage, 50, 50))
		if err != nil {
			log.Println(err)
			return ""
		}
		err = r.imageRepo.SaveImage(ctx, path, fileName, utils.ResizeImage(pngImage, 180, 180))
		if err != nil {
			return ""
		}
		return path + fileName
	}()

	updater := userModel.UserUpdater{
		UserID:      info.UserID,
		UserName:    info.DisplayName,
		Email:       info.Email,
		Birthday:    info.Birthday,
		Gender:      info.Gender,
		ProfilePath: profilePath,
		State:       userModel.UserStateActive,
		IsArtist:    &info.RegAsArtist,
	}

	err = r.userRepo.UpdateUserByAuthID(ctx, info.AuthID, &updater)
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
