//+build noasm
//+build !amd64
//+build appengine

package asm

import "github.com/rai-project/image"

func ResizeBilinear(inputImage image.RGBImage, height int, width int) (*image.RGBImage, error) {
	return nativeResizeBilinear(inputImage, height, width)
}
