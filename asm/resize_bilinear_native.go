package asm

import (
	"image"

	"github.com/bamiaux/rez"
	"github.com/disintegration/imaging"
	"github.com/pkg/errors"
	"github.com/c3sr/image/types"
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
	inputImage := toRGBAImage(srcPixels, srcWidth, srcHeight)
	resized := imaging.Resize(inputImage, targetWidth, targetHeight, imaging.Linear)
	res := &types.RGBImage{targetPixels, 3 * targetWidth, resized.Bounds()}
	res.FillFromNRGBAImage(resized)
	return nil
}

// not used
func resizeBilinearNativeOld(targetPixels []uint8, srcPixels []uint8, targetWidth, targetHeight, srcWidth, srcHeight int) error {
	inputImage := toRGBAImage(srcPixels, srcWidth, srcHeight)

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
