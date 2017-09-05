package image

import (
	"image"
	"io"

	"github.com/rai-project/go-libjpeg/jpeg"
)

func init() {
	libJpegDecoder := func(r io.Reader) (image.Image, error) {
		return jpeg.Decode(r, &jpeg.DecoderOptions{
			DisableFancyUpsampling: true,
		})
	}
	imageFormatDecoders["jpeg"] = libJpegDecoder
}
