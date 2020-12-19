package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	authCallback "pixstall-user/app/auth/delivery/model/auth-callback"
	getAuthURL "pixstall-user/app/auth/delivery/model/get-auth-url"
	"pixstall-user/domain/auth"
	authModel "pixstall-user/domain/auth/model"
)

type AuthController struct {
	authUseCase auth.UseCase
}

func NewAuthController(useCase auth.UseCase) AuthController {
	return AuthController{
		authUseCase: useCase,
	}
}

func (a AuthController) GetAuthURL(c *gin.Context) {
	authType := c.Query("type")
	if authType == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, getAuthURL.ErrorResponse{
			Message: "Auth type missing",
		})
		return
	}
	url, err := a.authUseCase.GetAuthURL(c, authType)
	if err != nil {
		c.AbortWithStatusJSON(getAuthURL.NewErrorResponse(err))
		return
	}
	c.JSON(getAuthURL.NewSuccessResponse(url))
}

func (a AuthController) AuthCallback(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")
	authType := c.Query("auth_type")
	if state == "" || code == "0" {
		return
	}
	handledAuthCallback, err := a.authUseCase.HandleAuthCallback(c, authModel.AuthCallback{
		AuthType: authType,
		State:    state,
		Code:     code,
	})
	if err != nil {
		c.AbortWithStatusJSON(authCallback.NewErrorResponse(err))
		return
	}
	c.PureJSON(authCallback.NewSuccessResponse(handledAuthCallback))
}