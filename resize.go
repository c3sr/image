package image

import (
	"image"
	"image/draw"
)

func AsRGBA(src image.Image) *image.RGBA {
	bounds := src.Bounds()
	img := image.NewRGBA(bounds)
	draw.Draw(img, bounds, src, bounds.Min, draw.Src)
	return img
}

// func Resize(ctx context.Context, inputImage image.Image, width, height int) (image.Image, error) {

// 	if span, newCtx := opentracing.StartSpanFromContext(ctx, "ImageResize"); span != nil {
// 		ctx = newCtx
// 		defer span.Finish()
// 	}

// 	inputImage = AsRGBA(inputImage)

// 	res := image.NewRGBA(image.Rectangle{
// 		Min: image.Point{
// 			X: 0,
// 			Y: 0,
// 		},
// 		Max: image.Point{
// 			X: width,
// 			Y: height,
// 		},
// 	})
// 	cfg, err := rez.PrepareConversion(res, inputImage)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "unable to create resize configuration")
// 	}
// 	converter, err := rez.NewConverter(cfg, rez.NewBilinearFilter())
// 	if err != nil {
// 		return nil, errors.Wrap(err, "unable to create resize converter")
// 	}
// 	err = converter.Convert(res, inputImage)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "unable to resize image")
// 	}

// 	return res, nil
// }
