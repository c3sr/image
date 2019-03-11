// +build cgo,!nolibjpeg

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
	"github.com/rai-project/image/types"
	"golang.org/x/image/bmp"
)

func jpegDecoder(mode types.Mode) func(io.Reader) (image.Image, error) {
	// if runtime.GOARCH == "ppc64le" {
	// 	return jpeg.Decode
	// }

	decodeOpts := &libjpeg.DecoderOptions{
		DCTMethod:              libjpeg.DCTISlow,
		DisableFancyUpsampling: true,
		// DisableBlockSmoothing:  true,
	}
	return func(r io.Reader) (image.Image, error) {
		if mode == types.RGBMode || mode == types.BGRMode {
			return libjpeg.DecodeIntoRGB(r, decodeOpts)
		}
		return libjpeg.Decode(r, decodeOpts)
	}
}

func getDecoder(format string, options *Options) (func(io.Reader) (image.Image, error), error) {
	if format == "jpeg" && options.resizeWidth != 0 && options.resizeHeight != 0 {
		if runtime.GOARCH == "ppc64le" {
			return jpeg.Decode, nil
		}
		return func(r io.Reader) (image.Image, error) {
			mode := options.mode
			decodeOpts := &libjpeg.DecoderOptions{
				ScaleTarget:            image.Rect(0, 0, options.resizeWidth, options.resizeHeight),
				DCTMethod:              libjpeg.DCTISlow,
				DisableFancyUpsampling: true,
				// DisableBlockSmoothing:  true,
			}
			if mode == types.RGBMode || mode == types.BGRMode {
				return libjpeg.DecodeIntoRGB(r, decodeOpts)
			}
			return libjpeg.Decode(r, decodeOpts)
		}, nil
	}
	imageFormatDecoders := map[string]func(io.Reader) (image.Image, error){
		"jpeg": jpegDecoder(options.mode),
		"png":  png.Decode,
		"gif":  gif.Decode,
		"bmp":  bmp.Decode,
	}
	if decoder, ok := imageFormatDecoders[format]; ok {
		return decoder, nil
	}
	return nil, errors.Errorf("format %v is not supported", format)
}
