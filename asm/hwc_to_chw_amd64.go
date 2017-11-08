//+build !noasm
//+build !appengine

package asm

import (
	"unsafe"

	"github.com/anthonynsimon/bild/parallel"
	"github.com/rai-project/cpu/cpuid"
)

//go:noescape
func __hwc2chw(result unsafe.Pointer, input unsafe.Pointer, height uint64, width uint64)

func Hwc2Chw(output []uint8, input []uint8, height int, width int) {
	if !cpuid.SupportsAVX() {
		nativeHwc2Chw(output, input, height, width)
		return
	}
	parallel.Line(height, func(start, end int) {
		offset := start * width * 3
		__hwc2chw(
			unsafe.Pointer(&output[offset]),
			unsafe.Pointer(&input[offset]),
			uint64(end-start),
			uint64(width),
		)
	})

}
