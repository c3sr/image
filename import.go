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
		return nil, errors.Errorf("unsuported image format. unable to find decoder %s in the image format decoder list", imageType)
	}

	buf := bytes.NewBuffer(image.Image())
	return decodeAsRGBImage(ctx, decoder, buf)
}

func decodeAsRGBImage(ctx context.Context, decoder func(io.Reader) (image.Image, error), reader io.Reader) (*RGBImage, error) {
	img, err := decoder(reader)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read %s as a jpeg image", ctx.Value("filePath"))
	}
	model := img.ColorModel()
	if model != color.RGBAModel {
		return nil, errors.New("expecting jpeg image to be in RGBA fromat")
	}
	rgbaImage, ok := img.(*image.RGBA)
	if !ok {
		return nil, errors.New("unable to cast to an rgba image")
	}

	return fromRGBAImage(ctx, rgbaImage)
}

func fromRGBAImage(ctx context.Context, rgbaImage *image.RGBA) (*RGBImage, error) {
	rgbImage := NewRGBImage(rgbaImage.Bounds())

	width := rgbaImage.Bounds().Dx()
	height := rgbaImage.Bounds().Dy()
	stride := rgbaImage.Stride

	rgbaImagePixels := rgbaImage.Pix
	rgbImagePixels := rgbImage.Pix
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rgbaOffset := y*stride + x*4
			rgbOffset := 3 * (y*width + x)
			rgbImagePixels[rgbOffset+0] = float32(rgbaImagePixels[rgbaOffset+0])
			rgbImagePixels[rgbOffset+1] = float32(rgbaImagePixels[rgbaOffset+1])
			rgbImagePixels[rgbOffset+2] = float32(rgbaImagePixels[rgbaOffset+2])
		}
	}

	return rgbImage, nil
}
