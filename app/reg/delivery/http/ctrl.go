package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"net/http"
	"pixstall-user/app/reg/delivery/model/register"
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
	images, err := getMultipartFormImages(c, "profile")
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

func getMultipartFormImages(ctx *gin.Context, key string) (*[]model3.ImageFile, error) {
	form, err := ctx.MultipartForm()
	if err != nil {
		return nil, err
	}
	fileHeaders := form.File[key]
	imageFiles := make([]model3.ImageFile, 0)
	for _, header := range fileHeaders {
		f, err := header.Open()
		if err != nil {
			continue
		}
		contentType, err := getFileContentType(f)
		if err != nil {
			_ = f.Close()
			continue
		}
		_, err = f.Seek(0, 0)
		if err != nil {
			_ = f.Close()
			continue
		}
		img, _, err := image.Decode(f)
		if err != nil {
			_ = f.Close()
			continue
		}
		_, err = f.Seek(0, 0)
		if err != nil {
			_ = f.Close()
			continue
		}
		imgF := model3.ImageFile{
			File: model3.File{
				File:        f,
				Name:        header.Filename,
				ContentType: contentType,
				Volume:      header.Size,
			},
			Size: model3.Size{
				Width:  float64(img.Bounds().Dx()),
				Height: float64(img.Bounds().Dy()),
				Unit:   "px",
			},
		}
		imageFiles = append(imageFiles, imgF)
		_ = f.Close()
	}
	if len(imageFiles) <= 0 {
		return nil, errors.New("")
	}
	return &imageFiles, nil
}

func getFileContentType(out multipart.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}