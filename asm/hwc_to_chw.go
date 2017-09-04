//+build noasm
//+build appengine

package asm

func Hwc2Chw(output []float32, input []float32, height int, width int) {
	nativeHwc2Chw(output, input, height, width)
}
