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
	"pixstall-user/domain/token"
	"pixstall-user/domain/user"
	userModel "pixstall-user/domain/user/model"
	"strings"
	"time"
)

type regUseCase struct {
	userRepo      user.Repo
	userMsgBroker user.MsgBroker
	imageRepo     domainImage.Repo
	tokenRepo     token.Repo
}

func NewRegUseCase(userRepo user.Repo, userMsgBroker user.MsgBroker, imageRepo domainImage.Repo, tokenRepo token.Repo) reg.UseCase {
	return &regUseCase{
		userRepo:      userRepo,
		userMsgBroker: userMsgBroker,
		imageRepo:     imageRepo,
		tokenRepo:     tokenRepo,
	}
}

func (r regUseCase) Registration(ctx context.Context, info model.RegInfo, pngImage image.Image) (*userModel.AuthUser, error) {

	//Check if authUser exist and in pending state
	extUser, err := r.userRepo.GetUserByAuthID(ctx, info.AuthID)
	if err != nil {
		return nil, err
	}
	if extUser.State != userModel.UserStatePending {
		return nil, model.RegErrorAuthIDAlreadyRegister
	}

	//Check if userId exist
	exist, err := r.userRepo.IsUserExist(ctx, info.UserID)
	if err != nil {
		return nil, err
	}
	if *exist {
		return nil, model.RegErrorDuplicateUser
	}

	//Upload image
	profilePath := r.uploadProfileImage(ctx, info.UserID, pngImage)

	info.ProfilePath = profilePath
	info.RegTime = time.Now()

	state := userModel.UserStateActive
	updater := userModel.UserUpdater{
		UserID:      info.UserID,
		UserName:    &info.DisplayName,
		Email:       &info.Email,
		Birthday:    &info.Birthday,
		Gender:      &info.Gender,
		ProfilePath: &profilePath,
		State:       &state,
		IsArtist:    &info.RegAsArtist,
		RegTime:     &info.RegTime,
	}

	err = r.userRepo.UpdateUserByAuthID(ctx, info.AuthID, &updater)
	if err != nil {
		return nil, model.RegErrorUnknown
	}

	if info.RegAsArtist {
		err := r.userMsgBroker.SendRegisterArtistMsg(ctx, &info)
		//not return err
		if err != nil {
			log.Printf("SendRegisterArtistMsg err %v", err)
		}
	} else {
		err := r.userMsgBroker.SendRegisterUserMsg(ctx, &info)
		//not return err
		if err != nil {
			log.Printf("SendRegisterUserMsg err %v", err)
		}
	}

	dUser, err := r.userRepo.GetUserDetails(ctx, info.UserID)
	if err != nil {
		log.Println(err)
		return nil, model.RegErrorUnknown
	}
	apiToken, err := r.tokenRepo.GenerateAPIToken(ctx, info.UserID)
	if err != nil {
		log.Println(err)
		return nil, model.RegErrorUnknown
	}

	dAuthUser := userModel.AuthUser{
		APIToken: apiToken,
		User:     *dUser,
	}

	return &dAuthUser, nil
}

func (r regUseCase) uploadProfileImage(ctx context.Context, userID string, pngImage image.Image) string {
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
	fileName = userID + "_" + fileName
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
}
