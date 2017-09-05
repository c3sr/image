package asm

import (
	goimage "image"

	"github.com/bamiaux/rez"
	"github.com/pkg/errors"
	"github.com/rai-project/image"
)

func nativeResizeBilinear(inputImage image.RGBImage, height int, width int) (*image.RGBImage, error) {
	res := image.NewRGBImage(goimage.Rect(0, 0, width, height))
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
