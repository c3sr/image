package asm

import "image"

func toRGBAImage(srcPixels []uint8, srcWidth, srcHeight int) *image.RGBA {
	inputImage := image.NewRGBA(image.Rect(0, 0, srcWidth, srcHeight))
	for ii := 0; ii < srcHeight; ii++ {
		inOffset := 3 * ii * srcWidth
		inputOffset := ii * inputImage.Stride
		for jj := 0; jj < srcWidth; jj++ {

			inputImage.Pix[inputOffset+0] = srcPixels[inOffset+0]
			inputImage.Pix[inputOffset+1] = srcPixels[inOffset+1]
			inputImage.Pix[inputOffset+2] = srcPixels[inOffset+2]
			inputImage.Pix[inputOffset+3] = 0xFF

			inOffset += 3
			inputOffset += 4
		}

	}

	return inputImage
}
