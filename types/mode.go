package types

// mode represents the image mode
type Mode int

const (
	RGBMode Mode = iota
	BGRMode
	NonInterlacedRGBMode
	NonInterlacedBGRMode
)
