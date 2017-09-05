//+build !noasm
//+build !appengine

package asm

import (
	"image"
	"unsafe"
)

//go:noescape
func __resize_vert(dst unsafe.Pointer, src unsafe.Pointer, dst_w uint64, dst_h uint64, src_h uint64)

//go:noescape
func __resize_hori(dst unsafe.Pointer, src unsafe.Pointer, dst_h uint64, dst_w uint64, src_w uint64)

//go:noescape
func __resize_bilinear(dst unsafe.Pointer, src unsafe.Pointer, dst_h uint64, dst_w uint64, src_h uint64, src_w uint64)

func ResizeBilinear(inputImage image.Image, height int, width int) (image.Image, error) {
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
	// __resize_bilinear(unsafe.Pointer(res.Pix), unsafe.Pointer(inputImage.RGBA.Pix)) {

	return res, nil
}
