//+build !noasm
//+build !appengine

package asm

import (
	"unsafe"

	"github.com/anthonynsimon/bild/parallel"
)

//go:noescape
func __Hwc2Cwh(output unsafe.Pointer, input unsafe.Pointer, height int, width int)

func Hwc2Cwh(output []float32, input []float32, width int, height int) {

	parallel.Line(height, func(start, end int) {
		offset := start * width * 3
		__Hwc2Cwh(unsafe.Pointer(&output[offset]), unsafe.Pointer(&input[offset]), width, end-start)
	})
}
