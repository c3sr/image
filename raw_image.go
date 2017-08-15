package image

import (
	"image"
	"image/color"
)

type Raw struct {
	Pix     []float32
	options Options
}

// ColorModel returns the Image's color model.
func (r Raw) ColorModel() color.Model {
	return RGBModel
}

func (r Raw) Mode() mode {
	return r.options.mode
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (r Raw) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: r.options.width,
			Y: r.options.height,
		},
	}
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (r Raw) At(x, y int) color.Color {
	offset := r.options.channels * (x + y*r.options.width)
	return RGB{
		R: r.Pix[offset+0],
		G: r.Pix[offset+1],
		B: r.Pix[offset+2],
	}
}
