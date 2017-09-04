//+build !noasm
//+build !appengine

package asm

import "unsafe"

//go:noescape
func __resize_vert(result unsafe.Pointer, input unsafe.Pointer, height uint64, width uint64)

//go:noescape
func __resize_hori(result unsafe.Pointer, input unsafe.Pointer, height uint64, width uint64)

//go:noescape
func __resize_bilinear(result unsafe.Pointer, input unsafe.Pointer, height uint64, width uint64)
