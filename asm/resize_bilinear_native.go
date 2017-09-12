package asm

import (
	"image"

	"github.com/bamiaux/rez"
	"github.com/pkg/errors"
	"github.com/rai-project/image/types"
)

func nativeResizeBilinear(in types.Image, targetHeight, targetWidth int) (types.Image, error) {

	srcHeight := in.Bounds().Dy()
	srcWidth := in.Bounds().Dx()

	switch in := in.(type) {
	case *types.RGBImage:
		res := types.NewRGBImage(image.Rect(0, 0, targetWidth, targetHeight))
		err := resizeBilinearNative(res.Pix, in.Pix, targetWidth, targetHeight, srcWidth, srcHeight)
		return res, err
	case *types.BGRImage:
		res := types.NewBGRImage(image.Rect(0, 0, targetWidth, targetHeight))
		err := resizeBilinearNative(res.Pix, in.Pix, targetWidth, targetHeight, srcWidth, srcHeight)
		return res, err
	}
	return nil, errors.New("invalid type while trying to resize image natively")
}

func resizeBilinearNative(targetPixels []uint8, srcPixels []uint8, targetWidth, targetHeight, srcWidth, srcHeight int) error {

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
