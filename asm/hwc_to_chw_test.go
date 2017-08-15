package asm

import (
	"math/rand"
	"testing"
)

func randomList(size int) []float32 {
	output := make([]float32, size)
	for ii := 0; ii < size; ii++ {
		output[ii] = rand.Float32()
	}
	return output
}

func benchmarkHwcToChw(b *testing.B, width, height, channels int) {
	input := randomList(width * height * channels)
	output := make([]float32, width*height*channels)
	for ii := 0; ii < b.N; ii++ {
		Hwc2Cwh(output, input, width, height)
	}
}

func benchmarkNativeHwcToChw(b *testing.B, width, height, channels int) {
	input := randomList(width * height * channels)
	output := make([]float32, width*height*channels)
	for ii := 0; ii < b.N; ii++ {
		nativeHwc2Cwh(output, input, width, height)
	}
}

func BenchmarkHwcToChw200x200(b *testing.B) {
	benchmarkHwcToChw(b, 200, 200, 3)
}

func BenchmarkNativeHwcToChw200x200(b *testing.B) {
	benchmarkNativeHwcToChw(b, 200, 200, 3)
}

func BenchmarkHwcToChw500x500(b *testing.B) {
	benchmarkHwcToChw(b, 500, 500, 3)
}

func BenchmarkNativeHwcToChw500x500(b *testing.B) {
	benchmarkNativeHwcToChw(b, 500, 500, 3)
}

func BenchmarkHwcToChw1000x1000(b *testing.B) {
	benchmarkHwcToChw(b, 1000, 1000, 3)
}

func BenchmarkNativeHwcToChw1000x1000(b *testing.B) {
	benchmarkNativeHwcToChw(b, 1000, 1000, 3)
}
