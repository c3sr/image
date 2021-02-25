package image

import (
	context "context"

	"github.com/c3sr/image/types"
)

type Options struct {
	ctx             context.Context
	resizeWidth     int
	resizeHeight    int
	resizeAlgorithm types.ResizeAlgorithm
	maxDimension    *int
	minDimension    *int
	keepAspectRatio *bool
	mode            types.Mode
	mean            [3]float32
	layout          Layout
	dctMethod       string
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
		o.resizeHeight = height
		o.resizeWidth = width
	}
}

func ResizeShape(shape []int) Option {
	if len(shape) != 2 {
		panic("expecting a resize shape of length 2 of the form height,width")
	}
	return func(o *Options) {
		o.resizeHeight = shape[0]
		o.resizeWidth = shape[1]
	}
}

func KeepAspectRatio(keepAspectRatio bool) Option {
	return func(o *Options) {
		o.keepAspectRatio = &keepAspectRatio
	}
}

func MaxDimension(dim int) Option {
	return func(o *Options) {
		o.maxDimension = &dim
	}
}

func MinDimension(dim int) Option {
	return func(o *Options) {
		o.minDimension = &dim
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
		maxDimension:    nil,
		minDimension:    nil,
		keepAspectRatio: nil,
		mode:            types.RGBMode,
		layout:          HWCLayout,
		dctMethod:       "INTEGER_ACCURATE",
		resizeAlgorithm: types.ResizeAlgorithmLinear,
		ctx:             context.Background(),
	}

	for _, o := range opts {
		o(options)
	}
	return options
}
