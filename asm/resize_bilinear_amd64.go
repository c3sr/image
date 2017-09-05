//+build !noasm
//+build !appengine

package asm

import (
	goimage "image"
	"unsafe"

	"github.com/anthonynsimon/bild/parallel"

	"github.com/rai-project/image"
)

//go:noescape
func __resize_bilinear(dst unsafe.Pointer, src unsafe.Pointer, dst_h uint64, dst_w uint64, src_h uint64, src_w uint64)

func ResizeBilinear(inputImage image.RGBImage, height int, width int) (*image.RGBImage, error) {
	res := image.NewRGBImage(goimage.Rect(0, 0, width, height))
	src_h := inputImage.Rect.Dy()
	src_w := inputImage.Rect.Dx()

	parallel.Line(src_h, func(start, end int) {
		offset := start * width * 3
		__resize_bilinear(
			unsafe.Pointer(&res.Pix[offset]),
			unsafe.Pointer(&inputImage.Pix[offset]),
			uint64(end-start), uint64(width),
			uint64(end-start), uint64(src_w),
		)
	})

	return res, nil
}
