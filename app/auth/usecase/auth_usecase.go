package usecase

import (
	"context"
	"pixstall-user/domain/auth"
	authModel "pixstall-user/domain/auth/model"
	"pixstall-user/domain/token"
	"pixstall-user/domain/user"
	"pixstall-user/domain/user/model"
)

type authUseCase struct {
	authRepo  auth.Repo
	userRepo  user.Repo
	tokenRepo token.Repo
}

func NewAuthUseCase(authRepo auth.Repo, userRepo user.Repo, tokenRepo token.Repo) auth.UseCase {
	return &authUseCase{
		authRepo: authRepo,
		userRepo: userRepo,
		tokenRepo: tokenRepo,
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
	user, err := a.userRepo.GetUserByAuthID(ctx, authUserInfo.ID)
	if err != nil {
		if userError, isError := err.(model.UserError); isError {
			switch userError {
			case model.UserErrorNotFound:
				user, err := a.createNewUser(ctx, authUserInfo)
				if err != nil {
					return nil, err
				}
				authToken, err := a.tokenRepo.GenerateAuthToken(ctx, user.UserID)
				if err != nil {
					return nil, err
				}
				return &authModel.HandledAuthCallback{
					Token: authToken,
					User:  *user,
				}, nil
			default:
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	if user.State == model.UserStateTerminated {
		return nil, model.UserErrorTerminated
	}

	//Existing User - generate new token
	authToken, err := a.tokenRepo.GenerateAuthToken(ctx, user.UserID)
	if err != nil {
		return nil, err
	}

	return &authModel.HandledAuthCallback{
		Token: authToken,
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
