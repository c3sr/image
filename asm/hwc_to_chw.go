//+build noasm
//+build appengine

package asm

func Hwc2Cwh(output []float32, input []float32, width int, height int) {
	nativeHwc2Cwh(output, input, width, height)
}
