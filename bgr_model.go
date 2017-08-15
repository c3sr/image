package image

import (
	"image/color"
)

var (
	BGRModel color.Model = color.ModelFunc(bgrModel)
)

// BGR represents a traditional 32-bit alpha-premultiplied color, having 8
// bits for each of red, green, blue.
type BGR struct {
	B, G, R float32
}

func (c BGR) RGBA() (r, g, b, a uint32) {
	r = uint32(255 * c.R)
	r |= r << 8
	g = uint32(255 * c.G)
	g |= g << 8
	b = uint32(255 * c.B)
	b |= b << 8
	a = uint32(0xFFFF)
	return
}

func bgrModel(c color.Color) color.Color {
	if _, ok := c.(BGR); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return BGR{float32(b >> 8), float32(g >> 8), float32(r >> 8)}
}
