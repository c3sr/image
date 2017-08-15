// +build !native

package image

import (
	"bytes"
	"image"
	"image/color"
	"io"

	context "golang.org/x/net/context"

	"github.com/Unknwon/com"
	"github.com/pkg/errors"
	bimg "gopkg.in/h2non/bimg.v1"
)

func read(ctx context.Context, filePath string, opts ...Option) (Image, error) {
	ctx = context.WithValue(ctx, "filePath", filePath)

	if !com.IsFile(filePath) {
		return nil, errors.Errorf("file %s not found while importing image", filePath)
	}
	options := NewOptions()
	for _, o := range opts {
		o(options)
	}

	ctx = context.WithValue(ctx, "options", options)

	buffer, err := bimg.Read(filePath)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read %s image file", filePath)
	}

	image := bimg.NewImage(buffer)
	imageSize, err := image.Size()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get image size")
	}

	processOptions := bimg.Options{}
	newWidth := options.width
	if newWidth == 0 {
		newWidth = imageSize.Width
	}
	newHeight := options.height
	if newHeight == 0 {
		newHeight = imageSize.Height
	}
	if newWidth != imageSize.Width || newHeight != imageSize.Height {
		processOptions.Width = newWidth
		processOptions.Height = newHeight
		processOptions.Force = true
		processOptions.Interpolator = bimg.Bilinear
	}

	processOptions.Type = bimg.PNG
	processOptions.Interpretation = bimg.InterpretationSRGB

	_, err = image.Process(processOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to process image")
	}

	imageType := image.Type()
	decoder, ok := imageFormatDecoders[imageType]
	if !ok {
		return nil, errors.Errorf("unsupported image format. "+
			"unable to find decoder %s in the image format decoder list", imageType)
	}

	buf := bytes.NewBuffer(image.Image())
	return decodeReader(ctx, decoder, buf)
}

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
