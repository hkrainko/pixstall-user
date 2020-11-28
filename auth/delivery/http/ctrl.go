package http

import (
	"github.com/gin-gonic/gin"
	get_auth_url "pixstall-user/auth/delivery/get-auth-url"
	"pixstall-user/domain/auth"
)

type AuthController struct{
	authUseCase auth.UseCase
}

func NewAuthController(usecase auth.UseCase) AuthController {
	return AuthController{
		authUseCase: usecase,
	}

}

func (a AuthController) GetAuthUrl(c *gin.Context) {
	authType := c.PostForm("authType")
	if authType == "" {
		return
	}
	url, err := a.authUseCase.GetAuthURL(c, authType)
	if err != nil {
		return
	}
	c.JSON(200, get_auth_url.NewResponse(url))
}