package usecase

import (
	"context"
	"log"
	domainFile "pixstall-user/domain/file"
	model2 "pixstall-user/domain/file/model"
	msgBroker "pixstall-user/domain/msg-broker"
	"pixstall-user/domain/reg"
	"pixstall-user/domain/reg/model"
	"pixstall-user/domain/token"
	"pixstall-user/domain/user"
	userModel "pixstall-user/domain/user/model"
	"time"
)

type regUseCase struct {
	userRepo  user.Repo
	msgBroker msgBroker.Repo
	fileRepo domainFile.Repo
	tokenRepo token.Repo
}

func NewRegUseCase(userRepo user.Repo, msgBroker msgBroker.Repo, fileRepo domainFile.Repo, tokenRepo token.Repo) reg.UseCase {
	return &regUseCase{
		userRepo:  userRepo,
		msgBroker: msgBroker,
		fileRepo: fileRepo,
		tokenRepo: tokenRepo,
	}
}

func (r regUseCase) Registration(ctx context.Context, info model.RegInfo, profile *model2.ImageFile) (*userModel.AuthUser, error) {

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
	var profilePath *string
	if profile != nil {
		profilePath, _ = r.fileRepo.SaveFile(ctx, profile.File, model2.FileTypeProfile, info.UserID, []string{"*"})
	}

	info.ProfilePath = profilePath
	info.RegTime = time.Now()

	state := userModel.UserStateActive
	updater := userModel.UserUpdater{
		UserID:      info.UserID,
		UserName:    &info.DisplayName,
		Email:       &info.Email,
		Birthday:    &info.Birthday,
		Gender:      &info.Gender,
		ProfilePath: profilePath,
		State:       &state,
		IsArtist:    &info.RegAsArtist,
		RegTime:     &info.RegTime,
	}

	err = r.userRepo.UpdateUserByAuthID(ctx, info.AuthID, &updater)
	if err != nil {
		return nil, model.RegErrorUnknown
	}

	if info.RegAsArtist {
		err := r.msgBroker.SendCreateArtistCmd(ctx, &info)
		//not return err
		if err != nil {
			log.Printf("SendCreateArtistCmd err %v", err)
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
