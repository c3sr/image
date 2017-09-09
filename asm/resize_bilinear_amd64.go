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

	err := IResizeBilinear(res.Pix, inputImage.Pix, width, height, srcWidth, srcHeight)

	return res, err
}

func IResizeBilinear(targetPixels []uint8, srcPixels []uint8, targetWidth, targetHeight, srcWidth, srcHeight int) error {

	parallel.Line(targetHeight, func(start, end int) {
		scale := float32(srcHeight) / float32(targetHeight)
		dstOffset := start * targetWidth * 3
		srcOffset := int(scale*float32(start)) * srcWidth * 3

		__resize_bilinear(
			unsafe.Pointer(&targetPixels[dstOffset]),
			unsafe.Pointer(&srcPixels[srcOffset]),
			uint64(end-start), uint64(targetWidth),
			uint64(scale*float32(end-start)), uint64(srcWidth),
		)
	})

	return nil
}
