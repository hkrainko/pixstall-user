package usecase

import (
	"context"
	"pixstall-user/domain/auth"
	authModel "pixstall-user/domain/auth/model"
	"pixstall-user/domain/reg"
	"pixstall-user/domain/token"
	"pixstall-user/domain/user"
	"pixstall-user/domain/user/model"
)

type authUseCase struct {
	authRepo  auth.Repo
	userRepo  user.Repo
	tokenRepo token.Repo
	regRepo   reg.Repo
}

func NewAuthUseCase(authRepo auth.Repo, userRepo user.Repo, tokenRepo token.Repo, regRepo reg.Repo) auth.UseCase {
	return &authUseCase{
		authRepo:  authRepo,
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
		regRepo:   regRepo,
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
	extUser, err := a.userRepo.GetUserByAuthID(ctx, authUserInfo.ID)
	if err != nil {
		if userError, isError := err.(model.UserError); isError {
			switch userError {
			case model.UserErrorNotFound:
				newUser, err := a.createNewUser(ctx, authUserInfo)
				if err != nil {
					return nil, err
				}
				regToken, err := a.regRepo.GenerateRegToken(ctx, newUser.AuthID)
				if err != nil {
					return nil, err
				}
				return &authModel.HandledAuthCallback{
					RegToken: regToken,
					User:     *newUser,
					PhotoUrl: authUserInfo.PhotoURL,
				}, nil
			default:
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	if extUser.State == model.UserStateTerminated {
		return nil, model.UserErrorTerminated
	}

	//Existing User - generate new token
	apiToken, err := a.tokenRepo.GenerateAPIToken(ctx, extUser.UserID)
	if err != nil {
		return nil, err
	}

	return &authModel.HandledAuthCallback{
		APIToken: apiToken,
		User:     *extUser,
		PhotoUrl: authUserInfo.PhotoURL,
	}, nil
}

func (a authUseCase) createNewUser(ctx context.Context, authUserInfo *authModel.AuthUserInfo) (*model.User, error) {

	user, err := a.userRepo.SaveAuthUser(ctx, authUserInfo)
	if err != nil {
		return nil, err
	}

	return user, nil
}
