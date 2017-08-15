package image

import (
	"image"
)

type Image interface {
	image.Image
}

func New(s string) (image.Image, error) {
	return nil, nil
}
