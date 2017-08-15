package image

type Options struct {
	width      int
	height     int
	channels   int
	interlaced bool
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

func Interlaced(bb bool) Option {
	return func(o *Options) {
		o.interlaced = bb
	}
}

func NewOptions() *Options {
	return &Options{}
}
