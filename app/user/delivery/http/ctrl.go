package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	get_user "pixstall-user/app/user/delivery/model/get-user"
	get_user_details "pixstall-user/app/user/delivery/model/get-user-details"
	update_user "pixstall-user/app/user/delivery/model/update-user"
	"pixstall-user/domain/user"
	"pixstall-user/domain/user/model"
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
	tokenUserID := c.GetString("userId")
	if tokenUserID == "" {
		c.JSON(get_user_details.NewErrorResponse(errors.New("userID format invalid")))
		return
	}
	dUser, err := u.userUseCase.GetAuthUser(c, tokenUserID)
	if err != nil {
		c.JSON(get_user_details.NewErrorResponse(err))
		return
	}
	c.JSON(get_user_details.NewSuccessResponse(dUser))
}

func (u UserController) UpdateUser(c *gin.Context) {
	pathUserID := c.Param("id")
	tokenUserID := c.GetString("userId")
	if tokenUserID == "" || pathUserID == "" {
		c.JSON(update_user.NewErrorResponse(errors.New("userID format invalid")))
		return
	}
	if tokenUserID != pathUserID {
		c.JSON(update_user.NewErrorResponse(errors.New("userIDs not match")))
		return
	}
	updater := model.UserUpdater{
		UserID: pathUserID,
	}
	if userName := c.PostForm("userName"); userName != "" {
		updater.UserName = &userName
	}
	if email := c.PostForm("email"); email != "" {
		updater.Email = &email
	}
	if birthday := c.PostForm("birthday"); birthday != "" {
		updater.Birthday = &birthday
	}
	if gender := c.PostForm("gender"); gender != "" {
		updater.Gender = &gender
	}

	dUser, err := u.userUseCase.UpdateUser(c, &updater)
	if err != nil {
		c.JSON(update_user.NewErrorResponse(err))
		return
	}
	c.JSON(update_user.NewSuccessResponse(dUser))
}
