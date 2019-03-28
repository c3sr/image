package image

import (
	context "context"

	"github.com/rai-project/image/types"
)

type Options struct {
	resizeWidth     int
	resizeHeight    int
	resizeAlgorithm types.ResizeAlgorithm
	mode            types.Mode
	mean            [3]float32
	layout          Layout
	dctMethod       string
	ctx             context.Context
}

type Option func(o *Options)

func Width(width int) Option {
	return func(o *Options) {
		o.resizeWidth = width
	}
}

func Height(height int) Option {
	return func(o *Options) {
		o.resizeHeight = height
	}
}

func Resized(height, width int) Option {
	return func(o *Options) {
		o.resizeWidth = width
		o.resizeHeight = height
	}
}

func ResizeAlgorithm(alg types.ResizeAlgorithm) Option {
	return func(o *Options) {
		o.resizeAlgorithm = alg
	}
}

func Mean(mean [3]float32) Option {
	return func(o *Options) {
		o.mean = mean
	}
}

func MeanValue(mean float32) Option {
	return Mean([3]float32{mean, mean, mean})
}

func ChannelLayout(layout Layout) Option {
	return func(o *Options) {
		o.layout = layout
	}
}

func Context(ctx context.Context) Option {
	return func(o *Options) {
		o.ctx = ctx
	}
}

func Mode(mode types.Mode) Option {
	return func(o *Options) {
		o.mode = mode
	}
}

func DCTMethod(method string) Option {
	return func(o *Options) {
		o.dctMethod = method
	}
}

func NewOptions(opts ...Option) *Options {
	options := &Options{
		mean:            [3]float32{0, 0, 0},
		mode:            types.RGBMode,
		layout:          HWCLayout,
		dctMethod:       "INTEGER_ACCURATE",
		resizeAlgorithm: types.ResizeAlgorithmBilinear,
		ctx:             context.Background(),
	}

	for _, o := range opts {
		o(options)
	}
	return options
}
