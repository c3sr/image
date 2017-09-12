package image

import (
	"image"

	"github.com/pkg/errors"
	"github.com/rai-project/image/asm"
	"github.com/rai-project/image/types"
)

func doResize(targetPixels []uint8, srcPixels []uint8, targetWidth, targetHeight, srcWidth, srcHeight int) error {
	return asm.IResizeBilinear(targetPixels, srcPixels, targetWidth, targetHeight, srcWidth, srcHeight)
}

func Resize(inputImage types.Image, opts ...Option) (types.Image, error) {
	options := NewOptions()
	for _, o := range opts {
		o(options)
	}

	width := options.resizeWidth
	height := options.resizeHeight

	switch in := inputImage.(type) {
	case *types.RGBImage:
		out := types.NewRGBImage(image.Rect(0, 0, width, height))
		inPix := in.Pix
		doResize(out.Pix, inPix, width, height, inputImage.Bounds().Dx(), inputImage.Bounds().Dy())
		return out, nil
	case *types.BGRImage:
		out := types.NewBGRImage(image.Rect(0, 0, width, height))
		inPix := in.Pix
		doResize(out.Pix, inPix, width, height, inputImage.Bounds().Dx(), inputImage.Bounds().Dy())
		return out, nil
	default:
		return nil, errors.New("input image was neither an RGB nor a BGR image")
	}

}
