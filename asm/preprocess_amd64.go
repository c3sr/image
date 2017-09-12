package asm

import (
	"github.com/pkg/errors"
	"github.com/rai-project/image/types"
)

func Preprocess(inputImage types.Image, height int, width int) (types.Image, error) {
	res, err := ResizeBilinear(inputImage, height, width)
	if err != nil {
		return nil, errors.Wrap(err, "unable to resize the image")
	}

	switch res := res.(type) {
	case *types.RGBImage:
		output := make([]uint8, len(res.Pix))
		Hwc2Chw(output, res.Pix, height, width)
		cnt := copy(res.Pix, output)
		if cnt != 3*height*width {
			return nil, errors.Wrap(err, "error copying after transposition")
		}
		return res, nil
	case *types.BGRImage:
		output := make([]uint8, len(res.Pix))
		Hwc2Chw(output, res.Pix, height, width)
		cnt := copy(res.Pix, output)
		if cnt != 3*height*width {
			return nil, errors.Wrap(err, "error copying after transposition")
		}
		return res, nil
	}

	return nil, errors.New("invalid input image")
}
