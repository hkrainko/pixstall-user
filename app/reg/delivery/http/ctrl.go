package http

import (
	"github.com/gin-gonic/gin"
	"image"
	"image/png"
	"pixstall-user/app/reg/delivery/model/register"
	model2 "pixstall-user/domain/artist/model"
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
		i, err := strconv.Atoi(c.PostForm("yearOfDrawing"))
		if err != nil {
			yearOfDrawing = i
		}
		artTypes = c.PostFormArray("artTypes")
	}
	profile, err := c.FormFile("profile")
	pngImage := func() image.Image {
		if err != nil {
			return nil
		}
		f, err := profile.Open()
		if err != nil {
			return nil
		}
		pngImage, err := png.Decode(f)
		if err != nil {
			return nil
		}
		return pngImage
	}()

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
			YearOfDrawing: &yearOfDrawing,
			ArtTypes:      &artTypes,
		},
	}
	authUser, err := r.regUseCase.Registration(c, regInfo, pngImage)
	if err != nil {
		c.JSON(register.NewErrorResponse(err))
		return
	}
	c.JSON(register.NewSuccessResponse(*authUser))
}
