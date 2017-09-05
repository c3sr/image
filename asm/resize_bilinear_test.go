package asm

import (
	goimage "image"
	"math/rand"
	"testing"

	"github.com/rai-project/image"
	"github.com/stretchr/testify/assert"
)

func randomRGBImage(h, w, c int) *image.RGBImage {
	res := image.NewRGBImage(goimage.Rect(0, 0, w, h))
	for ii := 0; ii < c*h*w; ii++ {
		res.Pix[ii] = rand.Float32()
	}
	return res
}

func randomRGBAImage(h, w, c int) *goimage.RGBA {
	res := goimage.NewRGBA(goimage.Rect(0, 0, w, h))
	for ii := 0; ii < c*h*w; ii++ {
		res.Pix[ii] = uint8(rand.Uint32())
	}
	return res
}

func benchmarkResizeBilinear(b *testing.B, height, width, channels int) {
	input := randomRGBImage(height, width, channels)
	for ii := 0; ii < b.N; ii++ {
		_, err := ResizeBilinear(input, height, width)
		assert.NoError(b, err)
	}
}

func benchmarkNativeResizeBilinear(b *testing.B, height, width, channels int) {
	input := randomRGBAImage(height, width, channels)
	for ii := 0; ii < b.N; ii++ {
		_, err := nativeResizeBilinear(input, height, width)
		assert.NoError(b, err)
	}
}

func BenchmarkResizeBilinear200x200(b *testing.B) {
	benchmarkResizeBilinear(b, 200, 200, 3)
}

func BenchmarkNativeResizeBilinear200x200(b *testing.B) {
	benchmarkNativeResizeBilinear(b, 200, 200, 3)
}

func BenchmarkResizeBilinear500x500(b *testing.B) {
	benchmarkResizeBilinear(b, 500, 500, 3)
}

func BenchmarkNativeResizeBilinear500x500(b *testing.B) {
	benchmarkNativeResizeBilinear(b, 500, 500, 3)
}

func BenchmarkResizeBilinear1000x1000(b *testing.B) {
	benchmarkResizeBilinear(b, 1000, 1000, 3)
}

func BenchmarkNativeResizeBilinear1000x1000(b *testing.B) {
	benchmarkNativeResizeBilinear(b, 1000, 1000, 3)
}
