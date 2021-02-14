// +build cgo,!nolibjpeg

package image

import (
	"image"
	"image/gif"
	"image/png"
	"io"
	"strings"

	"github.com/pkg/errors"
	libjpeg "github.com/c3sr/go-libjpeg"
	"github.com/c3sr/image/types"
	"golang.org/x/image/bmp"
)

func getDCTMethod(method string) (libjpeg.DCTMethod, error) {
	switch strings.ToLower(method) {
	case "slow", "integer_accurate", "dctislow":
		return libjpeg.DCTISlow, nil
	case "fast", "integer_fast", "dctifast":
		return libjpeg.DCTIFast, nil
	case "float", "dctfloat":
		return libjpeg.DCTFloat, nil
	default:
		var none libjpeg.DCTMethod
		return none, errors.Errorf("the DCT method %v specified is not valid", method)
	}
}

func jpegDecoder(options *Options) func(io.Reader) (image.Image, error) {
	return func(r io.Reader) (image.Image, error) {
		dctMethod, err := getDCTMethod(options.dctMethod)
		if err != nil {
			return nil, err
		}

		mode := options.mode

		decodeOpts := &libjpeg.DecoderOptions{
			DCTMethod:              dctMethod,
			DisableFancyUpsampling: true,
			// DisableBlockSmoothing:  true,
		}

		if mode == types.RGBMode || mode == types.BGRMode {
      res, err := libjpeg.DecodeIntoRGB(r, decodeOpts)
			return &types.RGBImage{res.Pix, res.Stride, res.Rectangle}
		}
		return libjpeg.Decode(r, decodeOpts)
	}
}

func getDecoder(format string, options *Options) (func(io.Reader) (image.Image, error), error) {
	if format == "jpeg" && options.resizeWidth != 0 && options.resizeHeight != 0 {
		// if runtime.GOARCH == "ppc64le" {
		// 	return jpeg.Decode, nil
		// }
		return func(r io.Reader) (image.Image, error) {
			mode := options.mode
			dctMethod, err := getDCTMethod(options.dctMethod)
			if err != nil {
				return nil, err
			}

			decodeOpts := &libjpeg.DecoderOptions{
				ScaleTarget:            image.Rect(0, 0, options.resizeWidth, options.resizeHeight),
				DCTMethod:              dctMethod,
				DisableFancyUpsampling: true,
				// DisableBlockSmoothing:  true,
			}
			if mode == types.RGBMode || mode == types.BGRMode {
        res, err := libjpeg.DecodeIntoRGB(r, decodeOpts)
        return &types.RGBImage{res.Pix, res.Stride, res.Rectangle}
			}
			return libjpeg.Decode(r, decodeOpts)
		}, nil
	}
	imageFormatDecoders := map[string]func(io.Reader) (image.Image, error){
		"jpeg": jpegDecoder(options),
		"png":  png.Decode,
		"gif":  gif.Decode,
		"bmp":  bmp.Decode,
	}
	if decoder, ok := imageFormatDecoders[format]; ok {
		return decoder, nil
	}
	return nil, errors.Errorf("format %v is not supported", format)
}
