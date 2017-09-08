package image

import (
	"bufio"
	"image"
	"io"

	"golang.org/x/net/context"

	"github.com/k0kubun/pp"
	"github.com/pkg/errors"
	"github.com/rai-project/image/types"
)

var ii = 0

func Read(ctx context.Context, reader0 io.Reader) (*types.RGBImage, error) {
	reader := bufio.NewReader(reader0)
	_, format, err := image.DecodeConfig(reader)
	if err != nil {
		return nil, err
	}
	reader.Reset(reader0)
	// format := "jpeg"
	// if ii == 1 {
	// 	format = "png"
	// }
	// ii++
	pp.Println(format)
	decoder, ok := imageFormatDecoders[format]
	if !ok {
		return nil, errors.Errorf("invalid format %v", format)
	}
	ctx = context.WithValue(ctx, "options", &Options{
		mode: types.RGBMode,
	})
	img, err := decodeReader(ctx, decoder, reader)
	if err != nil {
		return nil, errors.Wrap(err, "cannot decode reader")
	}
	res, ok := img.(*types.RGBImage)
	if !ok {
		return nil, errors.New("invalid return type for image read")
	}
	return res, nil
}
