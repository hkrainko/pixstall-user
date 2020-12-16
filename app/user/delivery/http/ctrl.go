package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	get_user "pixstall-user/app/user/delivery/model/get-user"
	"pixstall-user/domain/user"
)

type UserController struct {
	userUseCase user.UseCase
}

func NewUserController(useCase user.UseCase) UserController {
	return UserController{
		userUseCase: useCase,
	}
}

func (u UserController) GetUser(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(get_user.NewErrorResponse(errors.New("userID format invalid")))
		return
	}
	dUser, err := u.userUseCase.GetUser(c, userID)
	if err != nil {
		c.JSON(get_user.NewErrorResponse(err))
		return
	}
	c.JSON(get_user.NewSuccessResponse(dUser))
}
