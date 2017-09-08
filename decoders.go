package image

import (
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	context "golang.org/x/net/context"

	"github.com/k0kubun/pp"
	"github.com/pkg/errors"

	"github.com/rai-project/image/types"
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

func decodeReader(ctx context.Context, decoder func(io.Reader) (image.Image, error), reader io.Reader) (types.Image, error) {
	options, ok := ctx.Value("options").(*Options)
	if !ok {
		return nil, errors.New("expecting options to be passed in context")
	}
	if _, ok := ctx.Value("filePath").(string); !ok {
		ctx = context.WithValue(ctx, "filePath", "<<READER>>")
	}

	img, err := decoder(reader)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read %s as an image", ctx.Value("filePath"))
	}

	if res, ok := img.(*types.RGBImage); ok && options.mode == types.RGBMode {
		return res, nil
	}

	if res, ok := img.(*types.BGRImage); ok && options.mode == types.BGRMode {
		return res, nil
	}

	pp.Println(img)
	model := img.ColorModel()

	switch model {
	case types.RGBModel:
		if rgbImage, ok := img.(*types.RGBImage); ok {
			return rgbImage, nil
		}
		return nil, errors.New("unable to cast to an rgb image")
	case types.BGRModel:
		if bgrImage, ok := img.(*types.BGRImage); ok {
			return bgrImage, nil
		}
		return nil, errors.New("unable to cast to an bgr image")
	case color.NRGBAModel:
		if rgbaImage, ok := img.(*image.NRGBA); ok {
			return fromNRGBA(ctx, rgbaImage)
		}
		return nil, errors.New("unable to cast to an nrgba image")
	case color.RGBAModel:
		if rgbaImage, ok := img.(*image.RGBA); ok {
			return fromRGBA(ctx, rgbaImage)
		}
		return nil, errors.New("unable to cast to an rgba image")
	}

	return nil, errors.New("expecting image to be in RGBA or NRGBA fromat")
}

func fromNRGBA(ctx context.Context, nrgbaImage *image.NRGBA) (types.Image, error) {
	options, ok := ctx.Value("options").(*Options)
	if !ok {
		return nil, errors.New("expecting options to be passed in context")
	}
	switch options.mode {
	case types.RGBMode:
		rgbImage := types.NewRGBImage(nrgbaImage.Bounds())
		if err := rgbImage.FillFromNRGBAImage(ctx, nrgbaImage); err != nil {
			return nil, err
		}
		return rgbImage, nil
	case types.BGRMode:
		bgrImage := types.NewBGRImage(nrgbaImage.Bounds())
		if err := bgrImage.FillFromNRGBAImage(ctx, nrgbaImage); err != nil {
			return nil, err
		}
		return bgrImage, nil
	}
	return nil, errors.Errorf("invalid image mode %v", options.mode)
}

func fromRGBA(ctx context.Context, rgbaImage *image.RGBA) (types.Image, error) {
	options, ok := ctx.Value("options").(*Options)
	if !ok {
		return nil, errors.New("expecting options to be passed in context")
	}
	switch options.mode {
	case types.RGBMode:
		rgbImage := types.NewRGBImage(rgbaImage.Bounds())
		if err := rgbImage.FillFromRGBAImage(ctx, rgbaImage); err != nil {
			return nil, err
		}
		return rgbImage, nil
	case types.BGRMode:
		bgrImage := types.NewBGRImage(rgbaImage.Bounds())
		if err := bgrImage.FillFromRGBAImage(ctx, rgbaImage); err != nil {
			return nil, err
		}
		return bgrImage, nil
	}
	return nil, errors.Errorf("invalid image mode %v", options.mode)
}
