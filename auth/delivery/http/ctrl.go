package http

import (
	"github.com/gin-gonic/gin"
	authCallback "pixstall-user/auth/delivery/auth-callback"
	getAuthURL "pixstall-user/auth/delivery/get-auth-url"
	"pixstall-user/domain/auth"
	authModel "pixstall-user/domain/auth/model"
)

type AuthController struct{
	authUseCase auth.UseCase
}

func NewAuthController(usecase auth.UseCase) AuthController {
	return AuthController{
		authUseCase: usecase,
	}

}

func (a AuthController) GetAuthURL(c *gin.Context) {
	authType := c.PostForm("auth_type")
	if authType == "" {
		return
	}
	url, err := a.authUseCase.GetAuthURL(c, authType)
	if err != nil {
		return
	}
	c.JSON(200, getAuthURL.NewResponse(url))
}

func (a AuthController) AuthCallback(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")
	authType := c.Query("auth_type")
	if state == "" || code == "0" {
		return
	}
	userInfo, err := a.authUseCase.HandleAuthCallback(c, authModel.Callback{
		AuthType: authType,
		State:    state,
		Code:     code,
	})
	if err != nil {
		return
	}

	c.JSON(200, authCallback.NewResponse(userInfo))
}