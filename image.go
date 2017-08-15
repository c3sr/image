package image

import (
	"image"
)

type Image interface {
	image.Image
	Mode() mode
}

func New(s string) (Image, error) {
	return nil, nil
}
