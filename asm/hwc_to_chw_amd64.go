//+build !noasm
//+build !appengine

package c2goasmtest

import (
	"unsafe"
)

//go:noescape
func _HWC2CHW(output unsafe.Pointer, input unsafe.Pointer, height int, width int)

func HWC2CHW(output float32[], input float32[], width, height int ) {
	_HWC2CHW(output, input, width, height)
}
