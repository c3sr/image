package image

import (
	"image"
	"image/color"
)

// BGRImage is an in-memory image whose At method returns RGB values.
type BGRImage struct {
	// Pix holds the image's pixels, in R, G, B order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*3].
	Pix []float32
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func (p *BGRImage) ColorModel() color.Model { return BGRModel }

func (p *BGRImage) Bounds() image.Rectangle { return p.Rect }

func (p *BGRImage) Mode() mode { return BGRMode }

func (p *BGRImage) At(x, y int) color.Color {
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
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*4
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

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *BGRImage) Opaque() bool {
	if p.Rect.Empty() {
		return true
	}
	i0, i1 := 3, p.Rect.Dx()*4
	for y := p.Rect.Min.Y; y < p.Rect.Max.Y; y++ {
		for i := i0; i < i1; i += 4 {
			if p.Pix[i] != 0xff {
				return false
			}
		}
		i0 += p.Stride
		i1 += p.Stride
	}
	return true
}

// NewBGRImage returns a new BGRImage image with the given bounds.
func NewBGRImage(r image.Rectangle) *BGRImage {
	w, h := r.Dx(), r.Dy()
	buf := make([]float32, 3*w*h)
	return &BGRImage{buf, 3 * w, r}
}
