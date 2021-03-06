package types

import (
	"image"
	"image/color"

	"github.com/pkg/errors"
)

// RGBImage is an in-memory image whose At method returns RGB values.
type RGBImage struct {
	// Pix holds the image's pixels, in R, G, B order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*3].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func (p RGBImage) Channels() int { return 3 }

func (p RGBImage) Pixels() []uint8 { return p.Pix }

func (p RGBImage) ColorModel() color.Model { return RGBModel }

func (p RGBImage) Bounds() image.Rectangle { return p.Rect }

func (p RGBImage) Mode() Mode { return RGBMode }

func (p RGBImage) At(x, y int) color.Color {
	return p.RGBAt(x, y)
}

func (p *RGBImage) RGBAt(x, y int) RGB {
	if !(image.Point{x, y}.In(p.Rect)) {
		return RGB{}
	}
	i := p.PixOffset(x, y)
	return RGB{p.Pix[i+0], p.Pix[i+1], p.Pix[i+2]}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGBImage) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*3
}

func (p *RGBImage) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := RGBModel.Convert(c).(RGB)
	p.Pix[i+0] = c1.R
	p.Pix[i+1] = c1.G
	p.Pix[i+2] = c1.B
}

func (p *RGBImage) SetRGB(x, y int, c RGB) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.Pix[i+0] = c.R
	p.Pix[i+1] = c.G
	p.Pix[i+2] = c.B
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGBImage) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGBImage{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &RGBImage{
		Pix:    p.Pix[i:],
		Stride: p.Stride,
		Rect:   r,
	}
}

func (p *RGBImage) FillFromRGBImage(rgbImage *RGBImage) error {
	if p.Bounds() != rgbImage.Bounds() {
		return errors.Errorf("the bounds %v and %v did not match", p.Bounds(), rgbImage.Bounds())
	}

	copy(p.Pix, rgbImage.Pix)

	return nil
}

func (p *RGBImage) FillFromBGRImage(bgrImage *BGRImage) error {
	if p.Bounds() != bgrImage.Bounds() {
		return errors.Errorf("the bounds %v and %v did not match", p.Bounds(), bgrImage.Bounds())
	}

	width := bgrImage.Bounds().Dx()
	height := bgrImage.Bounds().Dy()
	stride := bgrImage.Stride

	bgrImagePixels := bgrImage.Pix
	rgbImagePixels := p.Pix
	for y := 0; y < height; y++ {
		rgbOffset := y * stride
		bgrOffset := y * p.Stride
		for x := 0; x < width; x++ {
			rgbImagePixels[rgbOffset+0] = bgrImagePixels[bgrOffset+2]
			rgbImagePixels[rgbOffset+1] = bgrImagePixels[bgrOffset+1]
			rgbImagePixels[rgbOffset+2] = bgrImagePixels[bgrOffset+0]
			rgbOffset += 3
			bgrOffset += 3
		}
	}

	return nil
}

func (p *RGBImage) ToRGBAImage() *image.RGBA {
	rgbaImage := image.NewRGBA(p.Bounds())

	width := rgbaImage.Bounds().Dx()
	height := rgbaImage.Bounds().Dy()
	stride := rgbaImage.Stride

	rgbaImagePixels := rgbaImage.Pix
	rgbImagePixels := p.Pix
	for y := 0; y < height; y++ {
		rgbaOffset := y * stride
		rgbOffset := y * p.Stride
		for x := 0; x < width; x++ {
			rgbaImagePixels[rgbaOffset+0] = rgbImagePixels[rgbOffset+0]
			rgbaImagePixels[rgbaOffset+1] = rgbImagePixels[rgbOffset+1]
			rgbaImagePixels[rgbaOffset+2] = rgbImagePixels[rgbOffset+2]
			rgbaImagePixels[rgbaOffset+3] = 0xFF
			rgbaOffset += 4
			rgbOffset += 3
		}
	}

	return rgbaImage
}

func (p *RGBImage) FillFromRGBAImage(rgbaImage *image.RGBA) error {
	if p.Bounds() != rgbaImage.Bounds() {
		return errors.Errorf("the bounds %v and %v did not match", p.Bounds(), rgbaImage.Bounds())
	}

	width := rgbaImage.Bounds().Dx()
	height := rgbaImage.Bounds().Dy()
	stride := rgbaImage.Stride

	rgbaImagePixels := rgbaImage.Pix
	rgbImagePixels := p.Pix
	for y := 0; y < height; y++ {
		rgbaOffset := y * stride
		rgbOffset := y * p.Stride
		for x := 0; x < width; x++ {
			rgbImagePixels[rgbOffset+0] = rgbaImagePixels[rgbaOffset+0]
			rgbImagePixels[rgbOffset+1] = rgbaImagePixels[rgbaOffset+1]
			rgbImagePixels[rgbOffset+2] = rgbaImagePixels[rgbaOffset+2]
			rgbaOffset += 4
			rgbOffset += 3
		}
	}

	return nil
}

func (p *RGBImage) FillFromGrayImage(grayImage *image.Gray) error {
	if p.Bounds() != grayImage.Bounds() {
		return errors.Errorf("the bounds %v and %v did not match", p.Bounds(), grayImage.Bounds())
	}

	width := grayImage.Bounds().Dx()
	height := grayImage.Bounds().Dy()
	stride := grayImage.Stride

	rgbImagePixels := p.Pix
	grayImagePixels := grayImage.Pix
	for y := 0; y < height; y++ {
		rgbOffset := y * p.Stride
		grayOffset := y * stride
		for x := 0; x < width; x++ {
			pix := grayImagePixels[grayOffset]
			rgbImagePixels[rgbOffset+0] = pix
			rgbImagePixels[rgbOffset+1] = pix
			rgbImagePixels[rgbOffset+2] = pix
			grayOffset++
			rgbOffset += 3
		}
	}

	return nil
}

func (p *RGBImage) FillFromNRGBAImage(nrgbaImage *image.NRGBA) error {
	if p.Bounds() != nrgbaImage.Bounds() {
		return errors.Errorf("the bounds %v and %v did not match", p.Bounds(), nrgbaImage.Bounds())
	}

	width := nrgbaImage.Bounds().Dx()
	height := nrgbaImage.Bounds().Dy()
	stride := nrgbaImage.Stride

	nrgbaImagePixels := nrgbaImage.Pix
	rgbImagePixels := p.Pix
	for y := 0; y < height; y++ {
		nrgbaOffset := y * stride
		rgbOffset := y * p.Stride
		for x := 0; x < width; x++ {
			rgbImagePixels[rgbOffset+0] = nrgbaImagePixels[nrgbaOffset+0]
			rgbImagePixels[rgbOffset+1] = nrgbaImagePixels[nrgbaOffset+1]
			rgbImagePixels[rgbOffset+2] = nrgbaImagePixels[nrgbaOffset+2]
			nrgbaOffset += 4
			rgbOffset += 3
		}
	}

	return nil
}

// NewRGBImage returns a new RGBImage image with the given bounds.
func NewRGBImage(r image.Rectangle) *RGBImage {
	w, h := r.Dx(), r.Dy()
	buf := make([]uint8, 3*w*h)
	return &RGBImage{buf, 3 * w, r}
}
