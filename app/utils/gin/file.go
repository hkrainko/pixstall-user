package gin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"image"
	"mime/multipart"
	"net/http"
	model2 "pixstall-user/domain/file/model"
)

func GetMultipartFormImages(ctx *gin.Context, key string) (*[]model2.ImageFile, error) {
	form, err := ctx.MultipartForm()
	if err != nil {
		return nil, err
	}
	fileHeaders := form.File[key]
	imageFiles := make([]model2.ImageFile, 0)
	for _, header := range fileHeaders {
		f, err := header.Open()
		if err != nil {
			continue
		}
		contentType, err := GetFileContentType(f)
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
		imgF := model2.ImageFile{
			File: model2.File{
				File:        f,
				Name:        header.Filename,
				ContentType: contentType,
				Volume:      header.Size,
			},
			Size: model2.Size{
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

func GetFileContentType(out multipart.File) (string, error) {

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