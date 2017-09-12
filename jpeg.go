// +build cgo

package image

import (
	"image"
	"image/gif"
	"image/png"
	"io"

	"github.com/pkg/errors"
	"github.com/rai-project/go-libjpeg"
	"golang.org/x/image/bmp"
)

func getDecoder(format string, options *Options) (func(io.Reader) (image.Image, error), error) {
	if format == "jpeg" && options.resizeWidth != 0 && options.resizeHeight != 0 {
		return func(r io.Reader) (image.Image, error) {
			return jpeg.Decode(r, &jpeg.DecoderOptions{
				ScaleTarget:            image.Rect(0, 0, options.resizeWidth, options.resizeHeight),
				DisableFancyUpsampling: true,
			})
		}, nil
	}
	imageFormatDecoders := map[string]func(io.Reader) (image.Image, error){
		"jpeg": func(r io.Reader) (image.Image, error) {
			return jpeg.Decode(r, &jpeg.DecoderOptions{
				DisableFancyUpsampling: true,
			})
		},
		"png": png.Decode,
		"gif": gif.Decode,
		"bmp": bmp.Decode,
	}
	if decoder, ok := imageFormatDecoders[format]; ok {
		return decoder, nil
	}
	return nil, errors.Errorf("format %v is not supported", format)
}
