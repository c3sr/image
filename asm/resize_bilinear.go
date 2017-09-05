//+build noasm
//+build appengine

package asm

import "image"

func ResizeBilinear(inputImage image.Image, height int, width int) (image.Image, error) {
	return nativeResizeBilinear(inputImage, height, width)
}
