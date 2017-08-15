package image

import (
	"image"
	"io"

	"github.com/pixiv/go-libjpeg/jpeg"
)

func init() {
	libJpegDecoder := func(r io.Reader) (image.Image, error) {
		return jpeg.Decode(r, &jpeg.DecoderOptions{
			DisableFancyUpsampling: true,
		})
	}
	imageFormatDecoders["jpeg"] = libJpegDecoder
}
