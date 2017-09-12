package asm

import (
	"image"

	"github.com/bamiaux/rez"
	"github.com/pkg/errors"
	"github.com/rai-project/image/types"
)

func nativeResizeBilinear(in *types.RGBImage, height int, width int) (*types.RGBImage, error) {
	inputImage := image.NewRGBA(in.Bounds())
	for ii := 0; ii < in.Bounds().Dy(); ii++ {
		for jj := 0; jj < in.Bounds().Dx(); jj++ {
			inOffset := ii*in.Stride + jj
			inputOffset := ii*inputImage.Stride + jj

			inputImage.Pix[inputOffset+0] = in.Pix[inOffset+0]
			inputImage.Pix[inputOffset+1] = in.Pix[inOffset+1]
			inputImage.Pix[inputOffset+2] = in.Pix[inOffset+2]
			inputImage.Pix[inputOffset+3] = 0xFF

			inOffset += 3
			inputOffset += 4
		}

	}

	tmp := image.NewRGBA(image.Rect(0, 0, width, height))
	cfg, err := rez.PrepareConversion(tmp, inputImage)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create resize configuration")
	}
	converter, err := rez.NewConverter(cfg, rez.NewBilinearFilter())
	if err != nil {
		return nil, errors.Wrap(err, "unable to create resize converter")
	}
	err = converter.Convert(tmp, inputImage)
	if err != nil {
		return nil, errors.Wrap(err, "unable to resize image")
	}

	res := types.NewRGBImage(image.Rect(0, 0, width, height))
	res.FillFromRGBAImage(tmp)

	return res, nil
}

func IResizeBilinearNative(targetPixels []uint8, srcPixels []uint8, targetWidth, targetHeight, srcWidth, srcHeight int) error {

	inputImage := image.NewRGBA(image.Rect(0, 0, srcWidth, srcHeight))
	for ii := 0; ii < srcHeight; ii++ {
		for jj := 0; jj < srcWidth; jj++ {
			inOffset := 3 * (ii*srcWidth + jj)
			inputOffset := ii*inputImage.Stride + 4*jj

			inputImage.Pix[inputOffset+0] = srcPixels[inOffset+0]
			inputImage.Pix[inputOffset+1] = srcPixels[inOffset+1]
			inputImage.Pix[inputOffset+2] = srcPixels[inOffset+2]
			inputImage.Pix[inputOffset+3] = 0xFF

			inOffset += 3
			inputOffset += 4
		}

	}

	tmp := image.NewRGBA(image.Rect(0, 0, targetWidth, targetHeight))
	cfg, err := rez.PrepareConversion(tmp, inputImage)
	if err != nil {
		return errors.Wrap(err, "unable to create resize configuration")
	}
	converter, err := rez.NewConverter(cfg, rez.NewBilinearFilter())
	if err != nil {
		return errors.Wrap(err, "unable to create resize converter")
	}
	err = converter.Convert(tmp, inputImage)
	if err != nil {
		return errors.Wrap(err, "unable to resize image")
	}

	res := &types.RGBImage{targetPixels, 3 * targetWidth, image.Rect(0, 0, targetWidth, targetHeight)}
	res.FillFromRGBAImage(tmp)

	return nil
}
