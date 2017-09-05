package asm

import (
	"github.com/pkg/errors"
	"github.com/rai-project/image"
)

func Preprocess(inputImage *image.RGBImage, height int, width int) (*image.RGBImage, error) {
	res, err := ResizeBilinear(inputImage, height, width)
	if err != nil {
		return nil, errors.Wrap(err, "unable to resize the image")
	}

	var output []float32
	Hwc2Chw(output, res.Pix, height, width)
	cnt := copy(res.Pix, output)
	if cnt != height*width {
		return nil, errors.Wrap(err, "error copying after transposition")
	}

	return res, nil
}
