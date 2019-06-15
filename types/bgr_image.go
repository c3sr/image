package types

import (
	"image"
	"image/color"

	"github.com/pkg/errors"
)

// BGRImage is an in-memory image whose At method returns RGB values.
type BGRImage struct {
	// Pix holds the image's pixels, in R, G, B order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*3].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func (p BGRImage) Channels() int { return 3 }

func (p BGRImage) Pixels() []uint8 { return p.Pix }

func (p BGRImage) ColorModel() color.Model { return BGRModel }

func (p BGRImage) Bounds() image.Rectangle { return p.Rect }

func (p BGRImage) Mode() Mode { return BGRMode }

func (p BGRImage) At(x, y int) color.Color {
	return p.BGRAt(x, y)
}

func (p *BGRImage) BGRAt(x, y int) BGR {
	if !(image.Point{x, y}.In(p.Rect)) {
		return BGR{}
	}
	i := p.PixOffset(x, y)
	return BGR{p.Pix[i+0], p.Pix[i+1], p.Pix[i+2]}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *BGRImage) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*3
}

func (p *BGRImage) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := RGBModel.Convert(c).(BGR)
	p.Pix[i+0] = c1.B
	p.Pix[i+1] = c1.G
	p.Pix[i+2] = c1.R
}

func (p *BGRImage) SetBGR(x, y int, c BGR) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.Pix[i+0] = c.B
	p.Pix[i+1] = c.G
	p.Pix[i+2] = c.R
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *BGRImage) SubImage(r image.Rectangle) Image {
	r = r.Intersect(p.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &BGRImage{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &BGRImage{
		Pix:    p.Pix[i:],
		Stride: p.Stride,
		Rect:   r,
	}
}

func (p *BGRImage) FillFromBGRImage(bgrImage *RGBImage) error {
	if p.Bounds() != bgrImage.Bounds() {
		return errors.Errorf("the bounds %v and %v did not match", p.Bounds(), bgrImage.Bounds())
	}

	copy(p.Pix, bgrImage.Pix)

	return nil
}

func (p *BGRImage) FillFromRGBImage(rgbImage *RGBImage) error {
	if p.Bounds() != rgbImage.Bounds() {
		return errors.Errorf("the bounds %v and %v did not match", p.Bounds(), rgbImage.Bounds())
	}

	width := rgbImage.Bounds().Dx()
	height := rgbImage.Bounds().Dy()
	stride := rgbImage.Stride

	rgbImagePixels := rgbImage.Pix
	bgrImagePixels := p.Pix
	for y := 0; y < height; y++ {
		rgbOffset := y * stride
		bgrOffset := y * p.Stride
		for x := 0; x < width; x++ {
			bgrImagePixels[bgrOffset+0] = rgbImagePixels[rgbOffset+2]
			bgrImagePixels[bgrOffset+1] = rgbImagePixels[rgbOffset+1]
			bgrImagePixels[bgrOffset+2] = rgbImagePixels[rgbOffset+0]
			rgbOffset += 3
			bgrOffset += 3
		}
	}

	return nil
}

func (p *BGRImage) ToRGBAImage() *image.RGBA {
	rgbaImage := image.NewRGBA(p.Bounds())

	width := rgbaImage.Bounds().Dx()
	height := rgbaImage.Bounds().Dy()
	stride := rgbaImage.Stride

	rgbaImagePixels := rgbaImage.Pix
	bgrImagePixels := p.Pix
	for y := 0; y < height; y++ {
		rgbaOffset := y * stride
		bgrOffset := y * p.Stride
		for x := 0; x < width; x++ {
			rgbaImagePixels[rgbaOffset+3] = 0xff
			rgbaImagePixels[rgbaOffset+2] = bgrImagePixels[bgrOffset+0]
			rgbaImagePixels[rgbaOffset+1] = bgrImagePixels[bgrOffset+1]
			rgbaImagePixels[rgbaOffset+0] = bgrImagePixels[bgrOffset+2]
			rgbaOffset += 4
			bgrOffset += 3
		}
	}

	return rgbaImage
}

func (p *BGRImage) FillFromRGBAImage(rgbaImage *image.RGBA) error {
	if p.Bounds() != rgbaImage.Bounds() {
		return errors.Errorf("the bounds %v and %v did not match", p.Bounds(), rgbaImage.Bounds())
	}

	width := rgbaImage.Bounds().Dx()
	height := rgbaImage.Bounds().Dy()
	stride := rgbaImage.Stride

	rgbaImagePixels := rgbaImage.Pix
	bgrImagePixels := p.Pix
	for y := 0; y < height; y++ {
		rgbaOffset := y * stride
		bgrOffset := y * p.Stride
		for x := 0; x < width; x++ {
			bgrImagePixels[bgrOffset+0] = rgbaImagePixels[rgbaOffset+2]
			bgrImagePixels[bgrOffset+1] = rgbaImagePixels[rgbaOffset+1]
			bgrImagePixels[bgrOffset+2] = rgbaImagePixels[rgbaOffset+0]
			rgbaOffset += 4
			bgrOffset += 3
		}
	}

	return nil
}

func (p *BGRImage) FillFromNRGBAImage(nrgbaImage *image.NRGBA) error {
	if p.Bounds() != nrgbaImage.Bounds() {
		return errors.Errorf("the bounds %v and %v did not match", p.Bounds(), nrgbaImage.Bounds())
	}

	width := nrgbaImage.Bounds().Dx()
	height := nrgbaImage.Bounds().Dy()
	stride := nrgbaImage.Stride

	nrgbaImagePixels := nrgbaImage.Pix
	bgrImagePixels := p.Pix
	for y := 0; y < height; y++ {
		nrgbaOffset := y * stride
		bgrOffset := y * p.Stride
		for x := 0; x < width; x++ {
			bgrImagePixels[bgrOffset+0] = nrgbaImagePixels[nrgbaOffset+2]
			bgrImagePixels[bgrOffset+1] = nrgbaImagePixels[nrgbaOffset+1]
			bgrImagePixels[bgrOffset+2] = nrgbaImagePixels[nrgbaOffset+0]
			nrgbaOffset += 4
			bgrOffset += 3
		}
	}

	return nil
}

func (p *BGRImage) FillFromGrayImage(grayImage *image.Gray) error {
	if p.Bounds() != grayImage.Bounds() {
		return errors.Errorf("the bounds %v and %v did not match", p.Bounds(), grayImage.Bounds())
	}

	width := grayImage.Bounds().Dx()
	height := grayImage.Bounds().Dy()
	stride := grayImage.Stride

	bgrImagePixels := p.Pix
	grayImagePixels := grayImage.Pix
	for y := 0; y < height; y++ {
		bgrOffset := y * p.Stride
		grayOffset := y * stride
		for x := 0; x < width; x++ {
			pix := grayImagePixels[grayOffset]
			bgrImagePixels[bgrOffset+0] = pix
			bgrImagePixels[bgrOffset+1] = pix
			bgrImagePixels[bgrOffset+2] = pix
			grayOffset++
			bgrOffset += 3
		}
	}

	return nil
}

// NewBGRImage returns a new BGRImage image with the given bounds.
func NewBGRImage(r image.Rectangle) *BGRImage {
	w, h := r.Dx(), r.Dy()
	buf := make([]uint8, 3*w*h)
	return &BGRImage{buf, 3 * w, r}
}
