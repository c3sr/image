// +build !cgo nolibjpeg

package image

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/pkg/errors"
	"golang.org/x/image/bmp"
)

func getDecoder(format string, options *Options) (func(io.Reader) (image.Image, error), error) {
	imageFormatDecoders := map[string]func(io.Reader) (image.Image, error){
		"jpeg": jpeg.Decode,
		"png":  png.Decode,
		"gif":  gif.Decode,
		"bmp":  bmp.Decode,
	}
	if decoder, ok := imageFormatDecoders[format]; ok {
		return decoder, nil
	}
	return nil, errors.Errorf("format %v is not supported", format)
}
