package image

type Options struct {
	width    int
	height   int
	channels int
	mode     mode
}

type Option func(o *Options)

func Width(ii int) Option {
	return func(o *Options) {
		o.width = ii
	}
}

func Height(ii int) Option {
	return func(o *Options) {
		o.height = ii
	}
}

func Channels(ii int) Option {
	return func(o *Options) {
		o.channels = ii
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
