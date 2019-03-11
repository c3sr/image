// +build cgo

package types

// #include "ConvertYCbCr.h"
import "C"

import (
	"image"
	"unsafe"

	"github.com/pkg/errors"
)

func (p *RGBImage) FillFromYCBCRImage(ycImage *image.YCbCr) error {
	if p.Bounds() != ycImage.Bounds() {
		return errors.Errorf("the bounds %v and %v did not match", p.Bounds(), ycImage.Bounds())
	}

	width := ycImage.Bounds().Dx()
	height := ycImage.Bounds().Dy()

	C.image_types_ImagingConvertYCbCr2RGB(
		(*C.uint8_t)(unsafe.Pointer(&p.Pix[0])),
		(*C.uint8_t)(unsafe.Pointer(&ycImage.Y[0])),
		(*C.uint8_t)(unsafe.Pointer(&ycImage.Cb[0])),
		(*C.uint8_t)(unsafe.Pointer(&ycImage.Cr[0])),
		C.int(ycImage.YStride),
		C.int(ycImage.CStride),
		C.int(width),
		C.int(height),
	)

	return nil
}

func (p *BGRImage) FillFromYCBCRImage(ycImage *image.YCbCr) error {
	if p.Bounds() != ycImage.Bounds() {
		return errors.Errorf("the bounds %v and %v did not match", p.Bounds(), ycImage.Bounds())
	}

	width := ycImage.Bounds().Dx()
	height := ycImage.Bounds().Dy()

	C.image_types_ImagingConvertYCbCr2BGR(
		(*C.uint8_t)(unsafe.Pointer(&p.Pix[0])),
		(*C.uint8_t)(unsafe.Pointer(&ycImage.Y[0])),
		(*C.uint8_t)(unsafe.Pointer(&ycImage.Cb[0])),
		(*C.uint8_t)(unsafe.Pointer(&ycImage.Cr[0])),
		C.int(ycImage.YStride),
		C.int(ycImage.CStride),
		C.int(width),
		C.int(height),
	)

	return nil
}
