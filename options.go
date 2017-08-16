package image

type Options struct {
	resizeWidth  int
	resizeHeight int
	mode         mode
}

type Option func(o *Options)

func Resize(width, height int) Option {
	return func(o *Options) {
		o.resizeWidth = width
		o.resizeHeight = height
	}
}

func Mean(mean [3]float32) Option {
	return func(o *Options) {
	}
}

func MeanValue(mean float32) Option {
	return func(o *Options) {
	}
}

func Layout(layout layout) Option {
	return func(o *Options) {
	}
}

func Mode(mode mode) Option {
	return func(o *Options) {
		o.mode = mode
	}
}

func NewOptions() *Options {
	return &Options{
		mode: RGBMode,
	}
}
