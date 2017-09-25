package image

import (
	"bufio"
	"io"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/rai-project/image/types"
)

// A reader is an io.Reader that can also peek ahead.
type reader interface {
	io.Reader
	Peek(int) ([]byte, error)
}

// asReader converts an io.Reader to a reader.
func asReader(r io.Reader) reader {
	if rr, ok := r.(reader); ok {
		return rr
	}
	return bufio.NewReader(r)
}

func getFormat(reader reader) (string, error) {
	formats := map[string]string{
		"jpeg": "\xff\xd8",
		"png":  "\x89PNG\r\n\x1a\n",
		"gif":  "GIF8?a",
		"bmp":  "BM????\x00\x00\x00\x00",
	}
	for format, magic := range formats {
		m, err := reader.Peek(len(magic))
		if err != nil {
			continue
		}
		if string(m) == magic {
			return format, nil
		}
	}
	return "", errors.New("input is not a valid image format")
}

func Read(r io.Reader, opts ...Option) (types.Image, error) {
	options := NewOptions(opts...)

	reader := asReader(r)
	format, err := getFormat(reader)
	if err != nil {
		return nil, err
	}
	decoder, err := getDecoder(format, options)
	if err != nil {
		return nil, err
	}

	span, ctx := options.tracer.StartSpanFromContext(options.ctx, "ReadImage", opentracing.Tags{"format": format})
	options.ctx = ctx
	defer span.Finish()

	img, err := decodeReader(decoder, reader, options)
	if err != nil {
		return nil, errors.Wrap(err, "cannot decode reader")
	}

	switch img.(type) {
	case *types.RGBImage, *types.BGRImage:
		if options.resizeWidth == 0 && options.resizeHeight == 0 {
			return img, nil
		}
		if img.Bounds().Dx() == options.resizeWidth && img.Bounds().Dy() == options.resizeHeight {
			return img, nil
		}
		return Resize(img, opts...)
	default:
		return nil, errors.New("invalid return type for image read")
	}

	return nil, errors.New("unreachable in image read")
}
