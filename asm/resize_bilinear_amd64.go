//+build !noasm
//+build !appengine

package asm

import (
	"unsafe"

	goimage "image"

	"github.com/anthonynsimon/bild/parallel"
	"github.com/rai-project/image/types"
)

//go:noescape
func __resize_bilinear(dst unsafe.Pointer, src unsafe.Pointer, dst_h uint64, dst_w uint64, src_h uint64, src_w uint64)

func ResizeBilinear(inputImage *types.RGBImage, height int, width int) (*types.RGBImage, error) {
	res := types.NewRGBImage(goimage.Rect(0, 0, width, height))
	srcHeight := inputImage.Rect.Dy()
	srcWidth := inputImage.Rect.Dx()

	parallel.Line(height, func(start, end int) {
		scale := float32(srcHeight) / float32(height)
		dstOffset := start * width * 3
		srcOffset := int(scale*float32(start)) * srcWidth * 3

		__resize_bilinear(
			unsafe.Pointer(&res.Pix[dstOffset]),
			unsafe.Pointer(&inputImage.Pix[srcOffset]),
			uint64(end-start), uint64(width),
			uint64(int(scale*float32(end))-int(scale*float32(start))), uint64(srcWidth),
		)
	})

	return res, nil
}
