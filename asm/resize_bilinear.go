//+build noasm
//+build !amd64
//+build appengine

package asm

import (
	"github.com/rai-project/image/types"
)

func ResizeBilinear(inputImage *types.RGBImage, height int, width int) (*types.RGBImage, error) {
	return nativeResizeBilinear(inputImage, height, width)
}
