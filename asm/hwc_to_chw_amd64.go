//+build !noasm
//+build !appengine

package asm

import (
	"unsafe"

	"github.com/anthonynsimon/bild/parallel"
)

//go:noescape
func __hwc2cwh(result unsafe.Pointer, input unsafe.Pointer, height uint64, width uint64)

//xxx go:noescape
// func ___hwc2cwh2(width uint32, height uint32, mean unsafe.Pointer, input unsafe.Pointer, result unsafe.Pointer)

func Hwc2Cwh(output []float32, input []float32, height int, width int) {
	// println("height  = ", uint32(height))
	// println("width  = ", uint32(width))

	parallel.Line(height, func(start, end int) {
		offset := start * width * 3
		__hwc2cwh(
			unsafe.Pointer(&output[offset]),
			unsafe.Pointer(&input[offset]),
			uint64(width),
			uint64(end-start),
		)
	})

}
