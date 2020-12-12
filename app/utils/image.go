package utils

import (
	"github.com/disintegration/imaging"
	"image"
)

func ResizeImage(originalImage image.Image, width int, height int) image.Image {
	return imaging.Fit(originalImage, width, height,imaging.Lanczos)
}