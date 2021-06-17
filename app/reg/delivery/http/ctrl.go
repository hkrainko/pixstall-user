package http

import (
	"github.com/gin-gonic/gin"
	_ "image/jpeg"
	_ "image/png"
	"pixstall-user/app/reg/delivery/model/register"
	gin2 "pixstall-user/app/utils/gin"
	model2 "pixstall-user/domain/artist/model"
	model3 "pixstall-user/domain/file/model"
	"pixstall-user/domain/reg"
	"pixstall-user/domain/reg/model"
	"strconv"
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
	authID := c.GetString("authId")
	if authID == "" {
		return
	}
	userID := c.PostForm("userId")
	displayName := c.PostForm("displayName")
	email := c.PostForm("email")
	birthday := c.PostForm("birthday")
	gender := c.PostForm("gender")
	regAsArtist := c.PostForm("regAsArtist")
	var yearOfDrawing int
	var artTypes []string
	if regAsArtist == "Y" {
		if yearOfDrawingStr, exist := c.GetPostForm("yearOfDrawing"); exist {
			i, err := strconv.Atoi(yearOfDrawingStr)
			if err == nil {
				yearOfDrawing = i
			}
		}
		artTypes = c.PostFormArray("artTypes")
	}
	var profileFile *model3.ImageFile
	images, err := gin2.GetMultipartFormImages(c, "profile")
	if err == nil {
		imgs := *images
		profileFile = &imgs[0]
	}

	if authID == "" {
		return
	}
	if userID == "" {
		return
	}
	regInfo := model.RegInfo{
		AuthID:      authID,
		UserID:      userID,
		DisplayName: displayName,
		Email:       email,
		Birthday:    birthday,
		Gender:      gender,
		RegAsArtist: func() bool {
			if regAsArtist == "Y" {
				return true
			}
			return false
		}(),
		RegArtistIntro: model2.ArtistIntro{
			YearOfDrawing: yearOfDrawing,
			ArtTypes:      artTypes,
		},
	}
	authUser, err := r.regUseCase.Registration(c, regInfo, profileFile)
	if err != nil {
		c.JSON(register.NewErrorResponse(err))
		return
	}
	c.JSON(register.NewSuccessResponse(*authUser))
}