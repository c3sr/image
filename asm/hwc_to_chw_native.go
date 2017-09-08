package asm

import (
	"github.com/anthonynsimon/bild/parallel"
)

func nativeHwc2Chw(output []uint8, input []uint8, height int, width int) {
	firstPlane := output[0*width*height:]
	secondPlane := output[1*width*height:]
	thirdPlane := output[2*width*height:]
	parallel.Line(height, func(start, end int) {
		w := width
		for y := start; y < end; y++ {
			for x := 0; x < width; x++ {
				offset := y*w + x
				firstPlane[offset] = input[3*offset+0]
				secondPlane[offset] = input[3*offset+1]
				thirdPlane[offset] = input[3*offset+2]
			}
		}
	})
}
