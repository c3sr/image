package image

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"golang.org/x/image/bmp"
)

var (
	imageFormatDecoders = map[string]func(io.Reader) (image.Image, error){
		"jpeg": jpeg.Decode,
		"png":  png.Decode,
		"gif":  gif.Decode,
		"bmp":  bmp.Decode,
	}
)
