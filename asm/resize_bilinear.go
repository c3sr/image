//+build noasm
//+build appengine

package asm

func ResizeBilinear(output []float32, input []float32, height int, width int) {
	nativeResizeBilinear(output, input, height, width)
}
