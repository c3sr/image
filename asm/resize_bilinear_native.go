package asm

import (
	"image"

	"github.com/bamiaux/rez"
	"github.com/pkg/errors"
	"github.com/rai-project/image/types"
)

func nativeResizeBilinear(inputImage *image.RGBAImage, height int, width int) (*image.RGBAImage, error) {
	res := types.NewRGBImage(image.Rect(0, 0, width, height))
	cfg, err := rez.PrepareConversion(res, inputImage)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create resize configuration")
	}
	converter, err := rez.NewConverter(cfg, rez.NewBilinearFilter())
	if err != nil {
		return nil, errors.Wrap(err, "unable to create resize converter")
	}
	err = converter.Convert(res, inputImage)
	if err != nil {
		return nil, errors.Wrap(err, "unable to resize image")
	}

	return res, nil
}
