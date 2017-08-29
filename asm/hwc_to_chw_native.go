package asm

import (
	"github.com/anthonynsimon/bild/parallel"
)

func nativeHwc2Cwh(output []float32, input []float32, height int, width int) {
	firstPlane := input[0*width*height:]
	secondPlane := input[1*width*height:]
	thirdPlane := input[2*width*height:]
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
