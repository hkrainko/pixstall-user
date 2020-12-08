package http

import (
	"github.com/gin-gonic/gin"
	model2 "pixstall-user/domain/artist/model"
	"pixstall-user/domain/reg"
	"pixstall-user/domain/reg/model"
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
	authID := c.PostForm("authId")
	userID := c.PostForm("userId")
	displayName := c.PostForm("displayName")
	email := c.PostForm("email")
	birthday := c.PostForm("birthday")
	gender := c.PostForm("gender")
	regAsArtist := c.PostForm("regAsArtist")

	if authID == "" {
		return
	}
	if userID == "" {
		return
	}
	regInfo := model.RegInfo{
		AuthID:        authID,
		UserID:        userID,
		DisplayName:   displayName,
		Email:         email,
		Birthday:      birthday,
		Gender:        gender,
		RegAsArtist:   func()bool { if regAsArtist == "Y" {return true}; return false}(),
		RegArtistInfo: model2.ArtistIntro{},
	}
	err := r.regUseCase.Registration(c, &regInfo)

	if err != nil {
		return
	}
}