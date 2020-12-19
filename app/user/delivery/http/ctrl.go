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

//For all user
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

//For authed user
func (u UserController) GetUserDetails(c *gin.Context) {
	pathUserID := c.Param("id")
	tokenUserID := c.GetString("userId")
	if tokenUserID == "" || pathUserID == "" {
		c.JSON(get_user.NewErrorResponse(errors.New("userID format invalid")))
		return
	}
	if tokenUserID != pathUserID {
		c.JSON(get_user.NewErrorResponse(errors.New("userIDs not match")))
		return
	}
	dUser, err := u.userUseCase.GetUserDetails(c, pathUserID)
	if err != nil {
		c.JSON(get_user.NewErrorResponse(err))
		return
	}
	c.JSON(get_user.NewSuccessResponse(dUser))
}