package image

// mode represents the image mode
type mode int

const (
	RGBMode mode = iota
	BGRMode
	NonInterlacedRGBMode
	NonInterlacedBGRMode
)
