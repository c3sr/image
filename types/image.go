package types

import (
	"image"
)

type Image interface {
	image.Image
	Mode() Mode
}
