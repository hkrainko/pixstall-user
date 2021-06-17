package usecase

import (
	"context"
	"pixstall-user/domain/file"
	model2 "pixstall-user/domain/file/model"
	"pixstall-user/domain/token"
	"pixstall-user/domain/user"
	"pixstall-user/domain/user/model"
)

type userUseCase struct {
	userRepo  user.Repo
	fileRepo  file.Repo
	tokenRepo token.Repo
}

func NewUserUseCase(userRepo user.Repo, fileRepo file.Repo, tokenRepo token.Repo) user.UseCase {
	return &userUseCase{
		userRepo:  userRepo,
		fileRepo:  fileRepo,
		tokenRepo: tokenRepo,
	}
}

func (u userUseCase) GetUser(ctx context.Context, userID string) (*model.User, error) {
	dUser, err := u.userRepo.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return dUser, nil
}

func (u userUseCase) GetAuthUser(ctx context.Context, userID string) (*model.AuthUser, error) {
	dUser, err := u.userRepo.GetUserDetails(ctx, userID)
	if err != nil {
		return nil, err
	}
	if dUser.State == model.UserStateTerminated {
		return nil, model.UserErrorTerminated
	}
	apiToken, err := u.tokenRepo.GenerateAPIToken(ctx, dUser.UserID)
	if err != nil {
		return nil, err
	}
	dUserDetails := model.AuthUser{
		APIToken: apiToken,
		User:     *dUser,
	}
	return &dUserDetails, nil
}

func (u userUseCase) UpdateUser(ctx context.Context, updater *model.UserUpdater, profile *model2.ImageFile) (*string, error) {
	dUser, err := u.userRepo.GetUserDetails(ctx, updater.UserID)
	if err != nil {
		return nil, err
	}
	if dUser.State == model.UserStateTerminated {
		return nil, model.UserErrorTerminated
	}
	var profilePath *string
	if profile != nil {
		profilePath, _ = u.fileRepo.SaveFile(ctx, profile.File, model2.FileTypeProfile, updater.UserID, []string{"*"})
	}
	updater.ProfilePath = profilePath
	err = u.userRepo.UpdateUser(ctx, updater.UserID, updater)
	if err != nil {
		return nil, err
	}

	return &updater.UserID, nil
}
