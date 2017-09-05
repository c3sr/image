//+build noasm
//+build appengine

package asm

import "github.com/rai-project/image"

func ResizeBilinear(inputImage image.RGBImage, height int, width int) (image.Image, error) {
	return nativeResizeBilinear(inputImage, height, width)
}
