package asm

import (
	"image"

	"github.com/disintegration/imaging"

	"github.com/pkg/errors"
	"github.com/c3sr/image/types"
)

func ResizeHermite(inputImage types.Image, height int, width int) (types.Image, error) {
	return nativeResizeHermite(inputImage, height, width)
}

func IResizeHermite(targetPixels []uint8, srcPixels []uint8, targetWidth, targetHeight, srcWidth, srcHeight int) error {
	return resizeHermiteNative(targetPixels, srcPixels, targetWidth, targetHeight, srcWidth, srcHeight)
}

func nativeResizeHermite(in types.Image, targetHeight, targetWidth int) (types.Image, error) {

	srcHeight := in.Bounds().Dy()
	srcWidth := in.Bounds().Dx()

	switch in := in.(type) {
	case *types.RGBImage:
		res := types.NewRGBImage(image.Rect(0, 0, targetWidth, targetHeight))
		err := resizeHermiteNative(res.Pix, in.Pix, targetWidth, targetHeight, srcWidth, srcHeight)
		return res, err
	case *types.BGRImage:
		res := types.NewBGRImage(image.Rect(0, 0, targetWidth, targetHeight))
		err := resizeHermiteNative(res.Pix, in.Pix, targetWidth, targetHeight, srcWidth, srcHeight)
		return res, err
	}
	return nil, errors.New("invalid type while trying to resize image natively")
}

func resizeHermiteNative(targetPixels []uint8, srcPixels []uint8, targetWidth, targetHeight, srcWidth, srcHeight int) error {
	inputImage := toRGBAImage(srcPixels, srcWidth, srcHeight)
	resized := imaging.Resize(inputImage, targetWidth, targetHeight, imaging.Hermite)
	res := &types.RGBImage{targetPixels, 3 * targetWidth, resized.Bounds()}
	res.FillFromNRGBAImage(resized)

	return nil
}
