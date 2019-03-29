package image

import (
	"image"

	"github.com/pkg/errors"
	"github.com/rai-project/image/asm"
	"github.com/rai-project/image/types"
	"github.com/rai-project/tracer"
)

func doResize(targetPixels []uint8, srcPixels []uint8, targetWidth, targetHeight, srcWidth, srcHeight int, resizeAlgorithm types.ResizeAlgorithm) error {
	switch resizeAlgorithm {
	case types.ResizeAlgorithmBiLinear:
		return asm.IResizeBilinear(targetPixels, srcPixels, targetWidth, targetHeight, srcWidth, srcHeight)
	case types.ResizeAlgorithmBiLinearASM:
		return asm.IResizeBilinearASM(targetPixels, srcPixels, targetWidth, targetHeight, srcWidth, srcHeight)
	case types.ResizeAlgorithmHermite:
		return asm.IResizeHermite(targetPixels, srcPixels, targetWidth, targetHeight, srcWidth, srcHeight)
	case types.ResizeAlgorithmNearestNeighbor:
		return asm.IResizeNearestNeighbor(targetPixels, srcPixels, targetWidth, targetHeight, srcWidth, srcHeight)
	}
	return errors.New("invalid resize algorithm")
}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func computeScaledDimension(inputImage types.Image, opts *Options) (int, int) {
	if opts.maxDimension == nil {
		return 0, 0
	}
	maxDimension := *opts.maxDimension
	imgWidth, imgHeight := inputImage.Bounds().Dx(), inputImage.Bounds().Dy()
	if opts.keepAspectRatio == nil || *opts.keepAspectRatio == false {
		return intMax(imgWidth, maxDimension), intMax(imgHeight, maxDimension)
	}

	resizeRatio := float32(maxDimension) / float32(intMax(imgWidth, imgHeight))
	width := int(resizeRatio * float32(imgWidth))
	height := int(resizeRatio * float32(imgHeight))

	return width, height
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

	if options.maxDimension != nil {
		targetWidth, targetHeight = computeScaledDimension(inputImage, options)
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
		doResize(out.Pix, inPix, targetWidth, targetHeight, srcWidth, srcHeight, options.resizeAlgorithm)
		return out, nil
	case *types.BGRImage:
		out := types.NewBGRImage(image.Rect(0, 0, targetWidth, targetHeight))
		inPix := in.Pix
		doResize(out.Pix, inPix, targetWidth, targetHeight, srcWidth, srcHeight, options.resizeAlgorithm)
		return out, nil
	default:
		return nil, errors.New("input image was neither an RGB nor a BGR image")
	}

}
