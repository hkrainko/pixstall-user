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

func (a authUseCase) HandleAuthCallback(ctx context.Context, authCallBack authModel.AuthCallback) (*authModel.HandledAuthCallback, error) {
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
				user, err := a.createNewUser(ctx, authUserInfo)
				if err != nil {
					return nil, user
				}
				return &authModel.HandledAuthCallback{
					Token: authUserInfo.Token,
					User:  *user,
				}, nil
			default:
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	//Existing User
	if user.State == model.UserStateInactive {
		return nil, model.UserErrorInactive
	}

	return &authModel.HandledAuthCallback{
		Token: authUserInfo.Token,
		User:  *user,
	}, nil
}

func (a authUseCase) createNewUser(ctx context.Context, authUserInfo *authModel.AuthUserInfo) (*model.User, error) {

	user, err := a.userRepo.SaveAuthUser(ctx, authUserInfo)
	if err != nil {
		return nil, err
	}

	return user, nil
}
