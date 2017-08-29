//+build noasm
//+build appengine

package asm

func Hwc2Cwh(output []float32, input []float32, height int, width int) {
	nativeHwc2Cwh(output, input, height, width)
}
