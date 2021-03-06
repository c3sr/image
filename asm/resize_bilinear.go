//+build !amd64 noasm appengine

package asm

import (
	"github.com/c3sr/image/types"
)

func ResizeBilinear(inputImage types.Image, height int, width int) (types.Image, error) {
	return nativeResizeBilinear(inputImage, height, width)
}

func IResizeBilinear(targetPixels []uint8, srcPixels []uint8, targetWidth, targetHeight, srcWidth, srcHeight int) error {
	return resizeBilinearNative(targetPixels, srcPixels, targetWidth, targetHeight, srcWidth, srcHeight)
}

func IResizeBilinearASM(targetPixels []uint8, srcPixels []uint8, targetWidth, targetHeight, srcWidth, srcHeight int) error {
	return resizeBilinearNative(targetPixels, srcPixels, targetWidth, targetHeight, srcWidth, srcHeight)
}
