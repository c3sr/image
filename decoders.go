package image

import (
	"image"
	"image/color"
	"io"

	"github.com/pkg/errors"

	"github.com/rai-project/image/types"
)

func decodeReader(decoder func(io.Reader) (image.Image, error), reader io.Reader, options *Options) (types.Image, error) {
	img, err := decoder(reader)
	if err != nil {
		return nil, errors.Wrap(err, "unable decode reader as image")
	}

	if res, ok := img.(*types.RGBImage); ok && options.mode == types.RGBMode {
		return res, nil
	}

	if res, ok := img.(*types.BGRImage); ok && options.mode == types.BGRMode {
		return res, nil
	}

	model := img.ColorModel()

	switch model {
	case types.RGBModel:
		if rgbImage, ok := img.(*types.RGBImage); ok {
			return fromRGB(rgbImage, options)
		}
		return nil, errors.New("unable to cast from an rgb image")
	case types.BGRModel:
		if bgrImage, ok := img.(*types.BGRImage); ok {
			return fromBGR(bgrImage, options)
		}
		return nil, errors.New("unable to cast from a bgr image")
	case color.RGBAModel:
		if rgbaImage, ok := img.(*image.RGBA); ok {
			return fromRGBA(rgbaImage, options)
		}
		return nil, errors.New("unable to cast from an rgba image")
	case color.NRGBAModel:
		if nrgbaImage, ok := img.(*image.NRGBA); ok {
			return fromNRGBA(nrgbaImage, options)
		}
		return nil, errors.New("unable to cast from an nrgba image")
	case color.GrayModel:
		if grayImage, ok := img.(*image.Gray); ok {
			return fromGray(grayImage, options)
		}
		return nil, errors.New("unable to cast from an nrgba image")
	case color.YCbCrModel:
		if ycbcrImage, ok := img.(*image.YCbCr); ok {
			return fromYCBCR(ycbcrImage, options)
		}
		return nil, errors.New("unable to cast from an ycbcr image")
	}

	return nil, errors.New("expecting image to be in RGBA or NRGBA fromat")
}

func fromGray(grayImage *image.Gray, options *Options) (types.Image, error) {

	switch options.mode {
	case types.RGBMode:
		rgbImage := types.NewRGBImage(grayImage.Bounds())
		if err := rgbImage.FillFromGrayImage(grayImage); err != nil {
			return nil, err
		}
		return rgbImage, nil
	case types.BGRMode:
		bgrImage := types.NewBGRImage(grayImage.Bounds())
		if err := bgrImage.FillFromGrayImage(grayImage); err != nil {
			return nil, err
		}
		return bgrImage, nil
	}
	return nil, errors.Errorf("invalid image mode %v", options.mode)
}

func fromNRGBA(nrgbaImage *image.NRGBA, options *Options) (types.Image, error) {
	switch options.mode {
	case types.RGBMode:
		rgbImage := types.NewRGBImage(nrgbaImage.Bounds())
		if err := rgbImage.FillFromNRGBAImage(nrgbaImage); err != nil {
			return nil, err
		}
		return rgbImage, nil
	case types.BGRMode:
		bgrImage := types.NewBGRImage(nrgbaImage.Bounds())
		if err := bgrImage.FillFromNRGBAImage(nrgbaImage); err != nil {
			return nil, err
		}
		return bgrImage, nil
	}
	return nil, errors.Errorf("invalid image mode %v", options.mode)
}

func fromYCBCR(ycbcrImage *image.YCbCr, options *Options) (types.Image, error) {
	switch options.mode {
	case types.RGBMode:
		rgbImage := types.NewRGBImage(ycbcrImage.Bounds())
		if err := rgbImage.FillFromYCBCRImage(ycbcrImage); err != nil {
			return nil, err
		}
		return rgbImage, nil
	case types.BGRMode:
		bgrImage := types.NewBGRImage(ycbcrImage.Bounds())
		if err := bgrImage.FillFromYCBCRImage(ycbcrImage); err != nil {
			return nil, err
		}
		return bgrImage, nil
	}
	return nil, errors.Errorf("invalid image mode %v", options.mode)
}

func fromRGB(rgbImage *types.RGBImage, options *Options) (types.Image, error) {
	switch options.mode {
	case types.RGBMode:
		return rgbImage, nil
	case types.BGRMode:
		res := types.NewBGRImage(rgbImage.Bounds())
		if err := res.FillFromRGBImage(rgbImage); err != nil {
			return nil, err
		}
		return res, nil
	}
	return nil, errors.Errorf("invalid image mode %v", options.mode)
}

func fromBGR(bgrImage *types.BGRImage, options *Options) (types.Image, error) {
	switch options.mode {
	case types.RGBMode:
		res := types.NewRGBImage(bgrImage.Bounds())
		if err := res.FillFromBGRImage(bgrImage); err != nil {
			return nil, err
		}
		return res, nil
	case types.BGRMode:
		return bgrImage, nil
	}
	return nil, errors.Errorf("invalid image mode %v", options.mode)
}

func fromRGBA(rgbaImage *image.RGBA, options *Options) (types.Image, error) {
	switch options.mode {
	case types.RGBMode:
		rgbImage := types.NewRGBImage(rgbaImage.Bounds())
		if err := rgbImage.FillFromRGBAImage(rgbaImage); err != nil {
			return nil, err
		}
		return rgbImage, nil
	case types.BGRMode:
		bgrImage := types.NewBGRImage(rgbaImage.Bounds())
		if err := bgrImage.FillFromRGBAImage(rgbaImage); err != nil {
			return nil, err
		}
		return bgrImage, nil
	}
	return nil, errors.Errorf("invalid image mode %v", options.mode)
}
