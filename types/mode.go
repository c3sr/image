package types

// mode represents the image mode
type Mode int

const (
	RGBMode Mode = iota
	BGRMode
	NonInterlacedRGBMode
	NonInterlacedBGRMode
	InvalidMode Mode = 9999
)

func (m Mode) Channels() int {
	switch m {
	case RGBMode, BGRMode, NonInterlacedBGRMode, NonInterlacedRGBMode:
		return 3
	case InvalidMode:
		return -1
	}
	return 0
}
