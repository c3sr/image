//+build !noasm
//+build !appengine

package asm

import (
	"unsafe"
)

//go:noescape
func __Hwc2Cwh(output unsafe.Pointer, input unsafe.Pointer, height int, width int)

func Hwc2Cwh(output []float32, input []float32, width int, height int) {
	__Hwc2Cwh(unsafe.Pointer(&output[0]), unsafe.Pointer(&input[0]), width, height)
}
