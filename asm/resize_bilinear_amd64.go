//+build !noasm
//+build amd64
//+build !appengine

package asm

import (
	"image"
	"unsafe"

	"github.com/anthonynsimon/bild/parallel"
	"github.com/pkg/errors"
	"github.com/c3sr/cpu/cpuid"
	"github.com/c3sr/image/types"
)

//go:noescape
func __resize_bilinear(dst unsafe.Pointer, src unsafe.Pointer, dst_h uint64, dst_w uint64, src_h uint64, src_w uint64)

func ResizeBilinear(in types.Image, targetHeight, targetWidth int) (types.Image, error) {

	if !cpuid.SupportsAVX() {
		return nativeResizeBilinear(in, targetHeight, targetWidth)
	}

	srcHeight := in.Bounds().Dy()
	srcWidth := in.Bounds().Dx()

	switch in := in.(type) {
	case *types.RGBImage:
		res := types.NewRGBImage(image.Rect(0, 0, targetWidth, targetHeight))
		err := IResizeBilinear(res.Pix, in.Pix, targetWidth, targetHeight, srcWidth, srcHeight)
		return res, err
	case *types.BGRImage:
		res := types.NewBGRImage(image.Rect(0, 0, targetWidth, targetHeight))
		err := IResizeBilinear(res.Pix, in.Pix, targetWidth, targetHeight, srcWidth, srcHeight)
		return res, err
	}
	return nil, errors.New("invalid type while trying to resize image natively")
}

func IResizeBilinear(targetPixels []uint8, srcPixels []uint8, targetWidth, targetHeight, srcWidth, srcHeight int) error {
	return resizeBilinearNative(targetPixels, srcPixels, targetWidth, targetHeight, srcWidth, srcHeight)
}

func IResizeBilinearASM(targetPixels []uint8, srcPixels []uint8, targetWidth, targetHeight, srcWidth, srcHeight int) error {

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
