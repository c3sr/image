package image

import (
	"image"

	"github.com/pkg/errors"
	"github.com/rai-project/image/asm"
	"github.com/rai-project/image/types"
	"github.com/rai-project/tracer"
)

func doResize(targetPixels []uint8, srcPixels []uint8, targetWidth, targetHeight, srcWidth, srcHeight int) error {
	return asm.IResizeBilinear(targetPixels, srcPixels, targetWidth, targetHeight, srcWidth, srcHeight)
}

func Resize(inputImage types.Image, opts ...Option) (types.Image, error) {
	options := NewOptions(opts...)

	srcWidth, srcHeight := inputImage.Bounds().Dx(), inputImage.Bounds().Dy()
	targetWidth, targetHeight := options.resizeWidth, options.resizeHeight

	if targetWidth == 0 {
		targetWidth = srcWidth
	}

	if targetHeight == 0 {
		targetHeight = srcHeight
	}

	if targetWidth == srcWidth && targetHeight == srcHeight {
		return inputImage, nil
	}

	if options.ctx != nil {
		if span, _ := tracer.StartSpanFromContext(options.ctx, tracer.APPLICATION_TRACE, "ResizeImage"); span != nil {
			span.SetTag("source_width", srcWidth)
			span.SetTag("source_height", srcHeight)
			span.SetTag("target_width", targetWidth)
			span.SetTag("target_height", targetHeight)
			defer span.Finish()
		}
	}

	switch in := inputImage.(type) {
	case *types.RGBImage:
		out := types.NewRGBImage(image.Rect(0, 0, targetWidth, targetHeight))
		inPix := in.Pix
		doResize(out.Pix, inPix, targetWidth, targetHeight, srcWidth, srcHeight)
		return out, nil
	case *types.BGRImage:
		out := types.NewBGRImage(image.Rect(0, 0, targetWidth, targetHeight))
		inPix := in.Pix
		doResize(out.Pix, inPix, targetWidth, targetHeight, srcWidth, srcHeight)
		return out, nil
	default:
		return nil, errors.New("input image was neither an RGB nor a BGR image")
	}

}
