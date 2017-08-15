package image

import (
	"image/color"
)

var (
	RGBModel color.Model = color.ModelFunc(rgbModel)
)

// RGB represents a traditional 32-bit alpha-premultiplied color, having 8
// bits for each of red, green, blue.
type RGB struct {
	R, G, B float32
}

func (c RGB) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = 0
	return
}

func rgbModel(c color.Color) color.Color {
	if _, ok := c.(RGB); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB{float32(r >> 8), float32(g >> 8), float32(b >> 8)}
}
