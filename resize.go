package image

import (
	"image"
	"math"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/c3sr/image/asm"
	"github.com/c3sr/image/types"
	"github.com/c3sr/tracer"
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

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func computeScaledDimension(inputImage types.Image, opts *Options) (int, int) {
	if opts.maxDimension == nil && opts.minDimension == nil {
		return 0, 0
	}

	if opts.maxDimension != nil && opts.minDimension != nil {
		return 0, 0
	}

	imgWidth, imgHeight := inputImage.Bounds().Dx(), inputImage.Bounds().Dy()

	var resizeRatio float32

	if opts.maxDimension != nil {
		maxDimension := *opts.maxDimension
		if opts.keepAspectRatio == nil || *opts.keepAspectRatio == false {
			return intMax(imgWidth, maxDimension), intMax(imgHeight, maxDimension)
		}
		resizeRatio = float32(maxDimension) / float32(intMax(imgWidth, imgHeight))
	} else {
		minDimension := *opts.minDimension
		if opts.keepAspectRatio == nil || *opts.keepAspectRatio == false {
			return intMin(imgWidth, minDimension), intMin(imgHeight, minDimension)
		}
		resizeRatio = float32(minDimension) / float32(intMin(imgWidth, imgHeight))
	}

	width := int(math.Round(float64(resizeRatio * float32(imgWidth))))
	height := int(math.Round(float64(resizeRatio * float32(imgHeight))))

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

	if options.maxDimension != nil || options.minDimension != nil {
		targetWidth, targetHeight = computeScaledDimension(inputImage, options)
	}

	if targetWidth == srcWidth && targetHeight == srcHeight {
		return inputImage, nil
	}

	if opentracing.SpanFromContext(options.ctx) != nil {
		if span, _ := tracer.StartSpanFromContext(options.ctx, tracer.APPLICATION_TRACE, "resize"); span != nil {
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
