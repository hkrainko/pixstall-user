package http

import (
	"github.com/gin-gonic/gin"
	"pixstall-user/app/domain/reg"
	"pixstall-user/app/domain/reg/model"
	userModel "pixstall-user/app/domain/user/model"
)

type RegController struct {
	regUseCase reg.UseCase
}

func NewRegController(useCase reg.UseCase) RegController {
	return RegController{
		regUseCase: useCase,
	}
}

func (r RegController) Registration(c *gin.Context) {
	userId := c.PostForm("userId")

	if userId == "" {
		return
	}
	regInfo := model.RegInfo{
		AuthID:        "",
		UserID:        userId,
		DisplayName:   "",
		Email:         "",
		Birthday:      "",
		Gender:        "",
		RegAsArtist:   false,
		RegArtistInfo: userModel.ArtistInfo{},
	}
	err := r.regUseCase.Registration(c, &regInfo)

	if err != nil {
		return
	}
}