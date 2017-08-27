//+build !noasm
//+build !appengine

package asm

import (
	"unsafe"
)

//go:noescape
func ___hwc2cwh(width uint32, height uint32, input unsafe.Pointer, result unsafe.Pointer)

//go:noescape
// func ___hwc2cwh2(width uint32, height uint32, mean unsafe.Pointer, input unsafe.Pointer, result unsafe.Pointer)

// func Hwc2Cwh(output []float32, input []float32, width int, height int) {

// 	parallel.Line(height, func(start, end int) {
// 		offset := start * width * 3
// 		__hwc2cwh((uint)width, (uint)end-start, unsafe.Pointer(&input[offset]), unsafe.Pointer(&output[offset]))
// 	})
// }
