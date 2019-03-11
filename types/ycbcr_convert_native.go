// +build !cgo

package types

import (
	"image"
	"image/color"

	"github.com/pkg/errors"
)

func (p *RGBImage) FillFromYCBCRImage(ycImage *image.YCbCr) error {
	if p.Bounds() != ycImage.Bounds() {
		return errors.Errorf("the bounds %v and %v did not match", p.Bounds(), ycImage.Bounds())
	}

	width := ycImage.Bounds().Dx()
	height := ycImage.Bounds().Dy()

	rgbImagePixels := p.Pix
	for y := 0; y < height; y++ {
		rgbOffset := y * p.Stride
		for x := 0; x < width; x++ {
			yi := ycImage.YOffset(x, y)
			ci := ycImage.COffset(x, y)
			r, g, b := color.YCbCrToRGB(
				ycImage.Y[yi],
				ycImage.Cb[ci],
				ycImage.Cr[ci],
			)
			rgbImagePixels[rgbOffset+0] = r
			rgbImagePixels[rgbOffset+1] = g
			rgbImagePixels[rgbOffset+2] = b
			rgbOffset += 3
		}
	}

	return nil
}

func (p *BGRImage) FillFromYCBCRImage(ycImage *image.YCbCr) error {
	if p.Bounds() != ycImage.Bounds() {
		return errors.Errorf("the bounds %v and %v did not match", p.Bounds(), ycImage.Bounds())
	}

	width := ycImage.Bounds().Dx()
	height := ycImage.Bounds().Dy()

	bgrImagePixels := p.Pix
	for y := 0; y < height; y++ {
		bgrOffset := y * p.Stride
		for x := 0; x < width; x++ {
			yi := ycImage.YOffset(x, y)
			ci := ycImage.COffset(x, y)
			r, g, b := color.YCbCrToRGB(
				ycImage.Y[yi],
				ycImage.Cb[ci],
				ycImage.Cr[ci],
			)
			bgrImagePixels[bgrOffset+0] = b
			bgrImagePixels[bgrOffset+1] = g
			bgrImagePixels[bgrOffset+2] = r
			bgrOffset += 3
		}
	}

	return nil
}
