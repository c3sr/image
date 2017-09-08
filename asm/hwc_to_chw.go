//+build noasm
//+build appengine

package asm

func Hwc2Chw(output []uint8, input []uint8, height int, width int) {
	nativeHwc2Chw(output, input, height, width)
}
