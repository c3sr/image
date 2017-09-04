package asm

import (
	"image"

	"github.com/bamiaux/rez"
	"github.com/pkg/errors"
)

func nativeResize(inputImage image.Image, width, height int) (image.Image, error) {
	res := image.NewRGBA(image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: width,
			Y: height,
		},
	})
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
