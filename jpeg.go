// +build cgo

package image

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"runtime"

	"github.com/pkg/errors"
	libjpeg "github.com/rai-project/go-libjpeg"
	"golang.org/x/image/bmp"
)

func jpegDecoder() func(io.Reader) (image.Image, error) {
	// if runtime.GOARCH == "ppc64le" {
	// 	return jpeg.Decode
	// }
	return func(r io.Reader) (image.Image, error) {
		return libjpeg.Decode(r, &libjpeg.DecoderOptions{
			DisableFancyUpsampling: true,
		})
	}
}

func getDecoder(format string, options *Options) (func(io.Reader) (image.Image, error), error) {
	if format == "jpeg" && options.resizeWidth != 0 && options.resizeHeight != 0 {
		if runtime.GOARCH == "ppc64le" {
			return jpeg.Decode, nil
		}
		return func(r io.Reader) (image.Image, error) {
			return libjpeg.Decode(r, &libjpeg.DecoderOptions{
				ScaleTarget:            image.Rect(0, 0, options.resizeWidth, options.resizeHeight),
				DisableFancyUpsampling: true,
			})
		}, nil
	}
	imageFormatDecoders := map[string]func(io.Reader) (image.Image, error){
		"jpeg": jpegDecoder(),
		"png":  png.Decode,
		"gif":  gif.Decode,
		"bmp":  bmp.Decode,
	}
	if decoder, ok := imageFormatDecoders[format]; ok {
		return decoder, nil
	}
	return nil, errors.Errorf("format %v is not supported", format)
}
