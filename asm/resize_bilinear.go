//+build !amd64

package asm

import (
	"github.com/rai-project/image/types"
)

func ResizeBilinear(inputImage types.Image, height int, width int) (types.Image, error) {
	return nativeResizeBilinear(inputImage, height, width)
}

func IResizeBilinear(targetPixels []uint8, srcPixels []uint8, targetWidth, targetHeight, srcWidth, srcHeight int) error {
	return resizeBilinearNative(targetPixels, srcPixels, targetWidth, targetHeight, srcWidth, srcHeight)
}
