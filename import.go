// +build !native

package image

import (
	"github.com/Unknwon/com"
	"github.com/k0kubun/pp"
	"github.com/pkg/errors"
	bimg "gopkg.in/h2non/bimg.v1"
)

func read(filePath string, opts ...Option) (*Image, error) {
	if !com.IsFile(filePath) {
		return nil, errors.Errorf("file %s not found while importing image", filePath)
	}
	options := NewOptions()
	for _, o := range opts {
		o(options)
	}
	buffer, err := bimg.Read(filePath)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read %s image file", filePath)
	}

	image := bimg.NewImage(buffer)
	imageSize, err := image.Size()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get image size")
	}

	processOptions := bimg.Options{}
	newWidth := options.width
	if newWidth == 0 {
		newWidth = imageSize.Width
	}
	newHeight := options.height
	if newHeight == 0 {
		newHeight = imageSize.Height
	}
	if newWidth != imageSize.Width || newHeight != imageSize.Height {
		processOptions.Width = newWidth
		processOptions.Height = newHeight
		processOptions.Force = true
	}
	if options.interlaced {
		processOptions.Interlace = options.interlaced
	}

	_, err = image.Process(processOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to process image")
	}

	pp.Println(len(image.Image()))

	return nil, nil

}
