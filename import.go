// +build !native

package image

// import (
// 	"bytes"

// 	context "golang.org/x/net/context"

// 	"github.com/Unknwon/com"
// 	"github.com/pkg/errors"
// 	bimg "gopkg.in/h2non/bimg.v1"
// )

// func read(ctx context.Context, filePath string, opts ...Option) (Image, error) {
// 	ctx = context.WithValue(ctx, "filePath", filePath)

// 	if !com.IsFile(filePath) {
// 		return nil, errors.Errorf("file %s not found while importing image", filePath)
// 	}
// 	options := NewOptions()
// 	for _, o := range opts {
// 		o(options)
// 	}

// 	ctx = context.WithValue(ctx, "options", options)

// 	buffer, err := bimg.Read(filePath)
// 	if err != nil {
// 		return nil, errors.Wrapf(err, "unable to read %s image file", filePath)
// 	}

// 	image := bimg.NewImage(buffer)
// 	imageSize, err := image.Size()
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to get image size")
// 	}

// 	processOptions := bimg.Options{}
// 	newWidth := options.width
// 	if newWidth == 0 {
// 		newWidth = imageSize.Width
// 	}
// 	newHeight := options.height
// 	if newHeight == 0 {
// 		newHeight = imageSize.Height
// 	}
// 	if newWidth != imageSize.Width || newHeight != imageSize.Height {
// 		processOptions.Width = newWidth
// 		processOptions.Height = newHeight
// 		processOptions.Force = true
// 		processOptions.Interpolator = bimg.Bilinear
// 	}

// 	processOptions.Type = bimg.PNG
// 	processOptions.Interpretation = bimg.InterpretationSRGB

// 	_, err = image.Process(processOptions)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to process image")
// 	}

// 	imageType := image.Type()
// 	decoder, ok := imageFormatDecoders[imageType]
// 	if !ok {
// 		return nil, errors.Errorf("unsupported image format. "+
// 			"unable to find decoder %s in the image format decoder list", imageType)
// 	}

// 	buf := bytes.NewBuffer(image.Image())
// 	return decodeReader(ctx, decoder, buf)
// }
