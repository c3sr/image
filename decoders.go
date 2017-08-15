package image

import (
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	context "golang.org/x/net/context"

	"github.com/pkg/errors"

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

func decodeReader(ctx context.Context, decoder func(io.Reader) (image.Image, error), reader io.Reader) (Image, error) {

	img, err := decoder(reader)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read %s as a jpeg image", ctx.Value("filePath"))
	}

	model := img.ColorModel()

	if model == color.NRGBAModel {
		if rgbaImage, ok := img.(*image.NRGBA); ok {
			return fromNRGBA(ctx, rgbaImage)
		}
		return nil, errors.New("unable to cast to an nrgba image")
	}

	if model == color.RGBAModel {
		if rgbaImage, ok := img.(*image.RGBA); ok {
			return fromRGBA(ctx, rgbaImage)
		}
		return nil, errors.New("unable to cast to an rgba image")
	}

	return nil, errors.New("expecting image to be in RGBA or NRGBA fromat")
}

func fromNRGBA(ctx context.Context, nrgbaImage *image.NRGBA) (Image, error) {
	options, ok := ctx.Value("options").(*Options)
	if !ok {
		return nil, errors.New("expecting options to be passed in context")
	}
	switch options.mode {
	case RGBMode:
		rgbImage := NewRGBImage(nrgbaImage.Bounds())
		if err := rgbImage.fillFromNRGBAImage(ctx, nrgbaImage); err != nil {
			return nil, err
		}
		return rgbImage, nil
	case BGRMode:
		bgrImage := NewBGRImage(nrgbaImage.Bounds())
		if err := bgrImage.fillFromNRGBAImage(ctx, nrgbaImage); err != nil {
			return nil, err
		}
		return bgrImage, nil
	}
	return nil, errors.Errorf("invalid image mode %v", options.mode)
}

func fromRGBA(ctx context.Context, rgbaImage *image.RGBA) (Image, error) {
	options, ok := ctx.Value("options").(*Options)
	if !ok {
		return nil, errors.New("expecting options to be passed in context")
	}
	switch options.mode {
	case RGBMode:
		rgbImage := NewRGBImage(rgbaImage.Bounds())
		if err := rgbImage.fillFromRGBAImage(ctx, rgbaImage); err != nil {
			return nil, err
		}
		return rgbImage, nil
	case BGRMode:
		bgrImage := NewBGRImage(rgbaImage.Bounds())
		if err := bgrImage.fillFromRGBAImage(ctx, rgbaImage); err != nil {
			return nil, err
		}
		return bgrImage, nil
	}
	return nil, errors.Errorf("invalid image mode %v", options.mode)
}
