package image

type Layout int

const (
	HWCLayout Layout = iota
	CHWLayout
	InvalidLayout Layout = 9999
)
