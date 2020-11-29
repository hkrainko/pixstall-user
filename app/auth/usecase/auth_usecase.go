package usecase

import (
	"context"
	"pixstall-user/app/domain/auth"
	authModel "pixstall-user/app/domain/auth/model"
	"pixstall-user/app/domain/user"
	"pixstall-user/app/domain/user/model"
)

type authUseCase struct {
	authRepo auth.Repo
	userRepo user.Repo
}

func NewAuthUseCase(authRepo auth.Repo, userRepo user.Repo) auth.UseCase {
	return &authUseCase{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}

func (a authUseCase) GetAuthURL(ctx context.Context, authType string) (string, error) {
	return a.authRepo.GetAuthURL(ctx, authType)
}

func (a authUseCase) HandleAuthCallback(ctx context.Context, authCallBack authModel.AuthCallback) (*authModel.AuthUserInfo, error) {
	authUserInfo, err := a.authRepo.GetAuthorizedUserInfo(ctx, authCallBack)
	if err != nil {
		return nil, err
	}

	//Get User
	user, err := a.userRepo.GetUser(ctx, authUserInfo.ID)
	if err != nil {
		if userError, isError := err.(model.UserError); isError {
			switch userError {
			case model.UserErrorNotFound:
				newUser, error := a.createNewUser(ctx, authUserInfo)
				if error != nil {
					return nil, error
				}

				return

				break
			default:
				return nil, err
			}
		} else {
			return nil, err
		}
	}


	//Existing User
	err = a.userRepo.SaveUser(ctx, user)
	if err != nil {
		return nil, err
	}
	
	return userInfo, nil
}

func (a authUseCase) createNewUser(ctx context.Context, authUserInfo *authModel.AuthUserInfo) (*model.User, error) {

	newUser := model.User{
		ID:       "",
		AuthID:   authUserInfo.ID,
		AuthType: authUserInfo.AuthType,
		Token:    "",
		Email:    authUserInfo.Email,
		Birthday: authUserInfo.Birthday,
		Gender:   authUserInfo.Gender,
		PhotoURL: authUserInfo.PhotoURL,
	}

	err := a.userRepo.SaveUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
