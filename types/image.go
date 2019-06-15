package types

import (
	"image"
)

type Image interface {
	image.Image
	Channels() int
	Pixels() []uint8
	Mode() Mode
}
