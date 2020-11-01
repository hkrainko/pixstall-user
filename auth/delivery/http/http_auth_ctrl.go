package http

import (
	"github.com/gin-gonic/gin"
	"pixstall_server/app/auth/delivery/get-auth-url"
	"pixstall_server/app/domain"
)

type AuthController struct{
	authUsecase domain.AuthUsecase
}

func NewAuthController(usecase domain.AuthUsecase) AuthController {
	return AuthController{
		authUsecase: usecase,
	}

}

func (a AuthController) GetAuthUrl(c *gin.Context) {
	authType := c.PostForm("authType")
	if authType == "" {
		return
	}
	url, err := a.authUsecase.GetAuthURL(c, authType)
	if err != nil {
		return
	}
	c.JSON(200, get_auth_url.NewResponse(url))
}